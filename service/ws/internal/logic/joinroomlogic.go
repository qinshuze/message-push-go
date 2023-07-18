package logic

import (
	"ccps.com/service/ws/internal/message"
	"context"

	"ccps.com/service/ws/internal/svc"
	"ccps.com/service/ws/ws"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinRoomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinRoomLogic {
	return &JoinRoomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// JoinRoom 加入房间
func (l *JoinRoomLogic) JoinRoom(in *ws.JoinRoomReq) (*ws.JoinRoomRes, error) {
	clientColl := message.Clients().GetBySelectOption(in.Options)

	for _, client := range clientColl.Slice() {
		for _, id := range in.RoomIds {
			client.JoinRoom(id)
		}
	}
	return &ws.JoinRoomRes{}, nil
}
