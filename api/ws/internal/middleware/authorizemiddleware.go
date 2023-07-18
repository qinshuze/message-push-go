package middleware

import (
	"ccps.com/api/ws/internal/config"
	"ccps.com/api/ws/internal/types"
	"ccps.com/internal/utils"
	"ccps.com/internal/utils/language"
	"ccps.com/internal/utils/response"
	"ccps.com/internal/utils/token"
	"encoding/json"
	"github.com/fatih/color"
	"net/http"
	"strings"
	"time"
)

type AuthorizeMiddleware struct {
	config.AuthConf
	*language.Loader
	authHandler func(info types.AuthInfo)
}

func NewAuthorizeMiddleware(conf config.AuthConf, langLoader *language.Loader, authHandler func(info types.AuthInfo)) *AuthorizeMiddleware {
	return &AuthorizeMiddleware{AuthConf: conf, Loader: langLoader, authHandler: authHandler}
}

func (m *AuthorizeMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			tokenStr = r.URL.Query().Get("token")
		} else {
			arr := strings.Split(tokenStr, " ")
			tokenStr = utils.ArrayItem(arr, 0, "")
		}

		if tokenStr == "" {
			color.Red("未认证的请求：%s", r.URL)
			response.NewJson(w).Error(response.ErrHttpUnauthorized)
			return
		}

		parse, err := token.Parse(m.AccessSecret, tokenStr)
		if err != nil {
			color.Red("token 解析错误：%s - %s", err.Error(), tokenStr)
			response.NewJson(w).Error(response.ErrTokenInvalid)
			return
		}

		var nowTime = utils.NowUnixSecond()
		if nowTime > parse.Expire {
			color.Red("token 过期：%s", time.Unix(parse.Expire, 0).Format("2006-01-02 15:04:05"))
			response.NewJson(w).Error(response.ErrTokenExpiration)
			return
		}

		payload := types.AuthInfo{}
		err = json.Unmarshal([]byte(parse.Payload), &payload)
		if err != nil {
			color.Red("token payload 解析错误：%s - %s", err.Error(), parse.Payload)
			response.NewJson(w).Error(response.ErrServerException)
			return
		}

		m.authHandler(payload)
		next(w, r)
	}
}
