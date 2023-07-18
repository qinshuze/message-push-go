package logic

import (
	"ccps.com/internal/utils"
	"ccps.com/service/ws/wsclient"
	"context"
	"github.com/fatih/color"
	"strings"
	"time"

	"ccps.com/api/ws/internal/svc"
	"ccps.com/api/ws/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClientListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClientListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientListLogic {
	return &ClientListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClientListLogic) ClientList(req *types.ClientListReq) (resp []types.ClientListRes, err error) {
	result, err := l.svcCtx.Ws.Clients(l.ctx, &wsclient.ClientsReq{
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

	startTime := time.Now().UnixNano()
	var list = make([]types.ClientListRes, 0)
	for _, client := range result.List {
		list = append(list, types.ClientListRes{
			Name:    client.Name,
			Tags:    client.Tags,
			RoomIds: utils.Nil(client.RoomIds, []string{}),
		})
	}

	color.Green("list: %v", time.Now().UnixNano()-startTime)

	return list, nil
}
