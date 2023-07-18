package logic

import (
	"ccps.com/service/ws/internal/message"
	"ccps.com/service/ws/internal/svc"
	"ccps.com/service/ws/ws"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type ClientsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClientsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientsLogic {
	return &ClientsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Clients 查询客户端
func (l *ClientsLogic) Clients(in *ws.ClientsReq) (*ws.ClientsRes, error) {
	clientColl := message.Clients().GetBySelectOption(in.Options)

	var list []*ws.Client
	for _, client := range clientColl.Slice() {
		list = append(list, &ws.Client{
			Id:      client.Id(),
			Name:    client.Name(),
			Tags:    client.Tags(),
			RoomIds: client.RoomIds(),
		})
	}

	return &ws.ClientsRes{List: list}, nil
}
