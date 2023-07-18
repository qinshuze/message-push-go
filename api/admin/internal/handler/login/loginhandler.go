package login

import (
	"ccps.com/api/admin/internal/logic/login"
	"ccps.com/api/admin/internal/svc"
	"ccps.com/api/admin/internal/types"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		_ = httpx.Parse(r, &req)
		if err := validator.New().Struct(req); err != nil {
			svcCtx.Response(w).Error(err)
			return
		}

		l := login.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			svcCtx.Response(w).Error(err)
			return
		}

		svcCtx.Response(w).R(resp)
	}
}
