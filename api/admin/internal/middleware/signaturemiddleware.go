package middleware

import (
	"ccps.com/api/admin/internal/lang"
	"ccps.com/internal/model"
	"ccps.com/internal/utils"
	"ccps.com/internal/utils/language"
	"ccps.com/internal/utils/response"
	"ccps.com/internal/utils/signature"
	"github.com/fatih/color"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"io"
	"net/http"
	"strconv"
)

type SignatureMiddleware struct {
	Db         sqlx.SqlConn
	LangLoader *language.Loader
}

func NewSignatureMiddleware(db sqlx.SqlConn, langLoader *language.Loader) *SignatureMiddleware {
	return &SignatureMiddleware{db, langLoader}
}

func (m *SignatureMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params = r.URL.Query()
		var accessKey = params.Get("access_key")
		var sign = params.Get("signature")
		var timestamp, err = strconv.ParseInt(params.Get("timestamp"), 10, 64)

		if accessKey == "" || sign == "" || err != nil {
			response.NewJson(w).Error(response.ErrSignatureInvalid)
			return
		}

		application, err := model.NewApplicationModel(m.Db).FindOneByAccessKey(r.Context(), accessKey)
		if err != nil {
			response.NewJson(w).
				SetStatusCode(http.StatusForbidden).E(
				response.CodeAccountNotFound,
				m.LangLoader.Trans(lang.AccountNotFound).Default("access_key is not ext"),
			)
			return
		}

		var hashedRequestPayload = params.Get("hashed_request_payload")
		var signStr = r.Method + "@" + r.Host + r.URL.Path + "?"
		if hashedRequestPayload != "" {
			bytes, err := io.ReadAll(r.Body)
			if err != nil {
				response.NewJson(w).Error(response.ErrHttpBadRequest)
				return
			}

			params.Set("hashed_request_payload", signature.Generate(application.AccessSecret, string(bytes)))
		}

		params.Del("signature")
		signStr += params.Encode()
		color.Green(signature.Generate(application.AccessSecret, signStr))
		if !signature.Equal(sign, signature.Generate(application.AccessSecret, signStr)) {
			response.NewJson(w).Error(response.ErrSignatureInvalid)
			return
		}

		//color.Red("current: %d, expire: %d, ss: %d", timestamp, utils.NowUnixSecond(), utils.NowUnixSecond()-timestamp)
		if utils.NowUnixSecond()-timestamp > 86400 {
			response.NewJson(w).Error(response.ErrSignatureExpiration)
			return
		}

		next(w, r)
	}
}
