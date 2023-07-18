package logic

import (
	"ccps.com/api/ws/internal/svc"
	"ccps.com/api/ws/internal/types"
	"ccps.com/internal/utils/response"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type ClientAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClientAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientAuthLogic {
	return &ClientAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClientAuthLogic) ClientAuth(req *types.ClientAuthReq) (resp *types.ClientAuthRes, err error) {
	if l.svcCtx.AuthInfo.AccessKey != req.Realm {
		return nil, response.ErrHttpForbidden
	}

	return &types.ClientAuthRes{}, nil
}
