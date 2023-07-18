package logic

import (
	"ccps.com/service/ws/internal/message"
	"ccps.com/service/ws/internal/svc"
	"ccps.com/service/ws/ws"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type CountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountLogic {
	return &CountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CountLogic) Count(in *ws.CountReq) (*ws.CountRes, error) {
	clientColl := message.Clients().GetBySelectOption(in.Options)

	return &ws.CountRes{Count: int64(clientColl.Count())}, nil
}
