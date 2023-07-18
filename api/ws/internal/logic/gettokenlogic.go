package logic

import (
	"ccps.com/api/ws/internal/svc"
	"ccps.com/api/ws/internal/types"
	"ccps.com/internal/utils/token"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenLogic {
	return &GetTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTokenLogic) GetToken(req *types.TokenReq) (resp *types.TokenRes, err error) {
	marshal, err := json.Marshal(types.AuthInfo{
		AccessKey: req.AccessKey,
	})
	if err != nil {
		return nil, err
	}

	tokenStr, err := token.Generate(l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessTTL, string(marshal))
	if err != nil {
		return nil, err
	}

	return &types.TokenRes{Token: tokenStr}, nil
}
