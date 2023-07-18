package logic

import (
	"ccps.com/service/ws/internal/message"
	"context"
	"github.com/fatih/color"

	"ccps.com/service/ws/internal/svc"
	"ccps.com/service/ws/ws"

	"github.com/zeromicro/go-zero/core/logx"
)

type PushLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPushLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushLogic {
	return &PushLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Push 推送消息
func (l *PushLogic) Push(in *ws.PushReq) (*ws.PushRes, error) {
	clientColl := message.Clients().GetBySelectOption(in.Options)

	var total, success, fail int64 = int64(clientColl.Count()), 0, 0
	for _, client := range clientColl.Slice() {
		err := client.Send(message.SendMessage{
			Type:    message.Relay,
			Content: in.Content,
			Sender:  client.Name(),
		})
		if err != nil {
			fail++
			color.Red("消息推送失败：%s - %s", client.Name(), err.Error())
		} else {
			success++
		}
	}
	return &ws.PushRes{TotalCount: total, SuccessCount: success, FailCount: fail}, nil
}
