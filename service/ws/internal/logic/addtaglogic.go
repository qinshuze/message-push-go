package logic

import (
	"ccps.com/service/ws/internal/message"
	"context"

	"ccps.com/service/ws/internal/svc"
	"ccps.com/service/ws/ws"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTagLogic {
	return &AddTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddTag 添加标签
func (l *AddTagLogic) AddTag(in *ws.AddTagsReq) (*ws.AddTagsResp, error) {
	clientColl := message.Clients().GetBySelectOption(in.Options)

	for _, client := range clientColl.Slice() {
		for _, tag := range in.Tags {
			client.AddTag(tag)
		}
	}

	return &ws.AddTagsResp{}, nil
}
