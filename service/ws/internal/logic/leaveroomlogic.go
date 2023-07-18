package logic

import (
	"ccps.com/service/ws/internal/message"
	"context"

	"ccps.com/service/ws/internal/svc"
	"ccps.com/service/ws/ws"

	"github.com/zeromicro/go-zero/core/logx"
)

type LeaveRoomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLeaveRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LeaveRoomLogic {
	return &LeaveRoomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// LeaveRoom 离开房间
func (l *LeaveRoomLogic) LeaveRoom(in *ws.LeaveRoomReq) (*ws.LeaveRoomRes, error) {
	clientColl := message.Clients().GetBySelectOption(in.Options)

	for _, client := range clientColl.Slice() {
		for _, id := range in.RoomIds {
			client.LeaveRoom(id)
		}
	}
	return &ws.LeaveRoomRes{}, nil
}
