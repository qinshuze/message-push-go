package logic

import (
	"ccps.com/internal/utils"
	"ccps.com/service/ws/wsclient"
	"context"
	"strings"

	"ccps.com/api/ws/internal/svc"
	"ccps.com/api/ws/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountClientLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCountClientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountClientLogic {
	return &CountClientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CountClientLogic) CountClient(req *types.CountClientReq) (resp *types.CountClientRes, err error) {
	count, err := l.svcCtx.Ws.Count(l.ctx, &wsclient.CountReq{
		Options: &wsclient.SelectOptions{
			Realms:  []string{l.svcCtx.AuthInfo.AccessKey},
			Names:   utils.FilterEmptyString(strings.Split(req.Names, ",")),
			Tags:    utils.FilterEmptyString(strings.Split(req.Tags, ",")),
			RoomIds: utils.FilterEmptyString(strings.Split(req.RoomIds, ",")),
		},
	})

	if err != nil {
		return nil, err
	}
	return &types.CountClientRes{Count: count.Count}, nil
}
