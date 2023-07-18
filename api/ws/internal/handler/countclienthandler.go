package handler

import (
	"net/http"

	"ccps.com/api/ws/internal/logic"
	"ccps.com/api/ws/internal/svc"
	"ccps.com/api/ws/internal/types"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CountClientHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CountClientReq
		_ = httpx.Parse(r, &req)
		if err := validator.New().Struct(req); err != nil {
			svcCtx.Response(w).Error(err)
			return
		}

		l := logic.NewCountClientLogic(r.Context(), svcCtx)
		resp, err := l.CountClient(&req)
		if err != nil {
			svcCtx.Response(w).Error(err)
			return
		}

		svcCtx.Response(w).R(resp)
	}
}
