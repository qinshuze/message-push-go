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

type LeaveRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLeaveRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LeaveRoomLogic {
	return &LeaveRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LeaveRoomLogic) LeaveRoom(req *types.LeaveRoomReq) (resp *types.LeaveRoomRes, err error) {
	_, err = l.svcCtx.Rpc.Ws.LeaveRoom(l.ctx, &wsclient.LeaveRoomReq{
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

	return &types.LeaveRoomRes{}, nil
}
