package handler

import (
	"net/http"

	"ccps.com/api/ws/internal/logic"
	"ccps.com/api/ws/internal/svc"
	"ccps.com/api/ws/internal/types"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ClientListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ClientListReq
		_ = httpx.Parse(r, &req)
		if err := validator.New().Struct(req); err != nil {
			svcCtx.Response(w).Error(err)
			return
		}

		l := logic.NewClientListLogic(r.Context(), svcCtx)
		resp, err := l.ClientList(&req)
		if err != nil {
			svcCtx.Response(w).Error(err)
			return
		}

		svcCtx.Response(w).R(resp)
	}
}
