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

type JoinRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJoinRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinRoomLogic {
	return &JoinRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JoinRoomLogic) JoinRoom(req *types.JoinRoomReq) (resp *types.JoinRoomRes, err error) {
	_, err = l.svcCtx.Rpc.Ws.JoinRoom(l.ctx, &wsclient.JoinRoomReq{
		Options: &wsclient.SelectOptions{
			Realms:  []string{l.svcCtx.AuthInfo.AccessKey},
			Names:   utils.FilterEmptyString(strings.Split(req.OptNames, ",")),
			RoomIds: utils.FilterEmptyString(strings.Split(req.OptRoomIds, ",")),
			Tags:    utils.FilterEmptyString(strings.Split(req.OptTags, ",")),
		},
		RoomIds: utils.FilterEmptyString(strings.Split(req.RoomIds, ",")),
	})

	if err != nil {
		return nil, err
	}

	return &types.JoinRoomRes{}, nil
}
