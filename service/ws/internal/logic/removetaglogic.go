package logic

import (
	"ccps.com/service/ws/internal/message"
	"context"

	"ccps.com/service/ws/internal/svc"
	"ccps.com/service/ws/ws"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveTagLogic {
	return &RemoveTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RemoveTag 移除标签
func (l *RemoveTagLogic) RemoveTag(in *ws.RemoveTagsReq) (*ws.RemoveTagsResp, error) {
	clientColl := message.Clients().GetBySelectOption(in.Options)

	for _, client := range clientColl.Slice() {
		for _, tag := range in.Tags {
			client.RemoveTag(tag)
		}
	}
	return &ws.RemoveTagsResp{}, nil
}
