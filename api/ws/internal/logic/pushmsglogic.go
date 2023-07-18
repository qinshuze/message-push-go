package logic

import (
	"ccps.com/api/ws/internal/svc"
	"ccps.com/api/ws/internal/types"
	"ccps.com/internal/utils"
	"ccps.com/service/ws/wsclient"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
)

type PushMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPushMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushMsgLogic {
	return &PushMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PushMsgLogic) PushMsg(req *types.PushMsgReq) (resp *types.PushMsgRes, err error) {
	result, err := l.svcCtx.Ws.Push(l.ctx, &wsclient.PushReq{
		Options: &wsclient.SelectOptions{
			Realms:  []string{l.svcCtx.AuthInfo.AccessKey},
			Names:   utils.FilterEmptyString(strings.Split(req.Names, ",")),
			Tags:    utils.FilterEmptyString(strings.Split(req.Tags, ",")),
			RoomIds: utils.FilterEmptyString(strings.Split(req.RoomIds, ",")),
		},
		Content: req.Content,
	})

	return &types.PushMsgRes{
		TotalCount:   int(result.TotalCount),
		SuccessCount: int(result.SuccessCount),
		FailCount:    int(result.FailCount),
	}, nil
}
