package login

import (
	"ccps.com/api/admin/internal/svc"
	"ccps.com/api/admin/internal/types"
	"ccps.com/internal/model"
	"ccps.com/internal/utils"
	"ccps.com/internal/utils/response"
	"ccps.com/internal/utils/token"
	"context"
	"encoding/json"
	"github.com/fatih/color"
	"golang.org/x/crypto/bcrypt"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	resp = &types.LoginRes{}
	userModel := model.NewUserModel(l.svcCtx.Db)

	// 查询登陆用户信息
	user, err := userModel.FindOneByAccount(l.ctx, req.Account)
	if err != nil {
		return nil, response.ErrAccountNotExist
	}

	// 验证密码是否正确
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, response.ErrAccountPwdError
	}

	// 获取用户角色
	roleRows, _, err := model.NewRoleModel(l.svcCtx.Db).FindByUserIds(l.ctx, []int64{user.Id})
	if err != nil {
		return nil, err
	}

	// 组装角色信息
	var roleIds []int64
	roles := make([]*types.Role, 0)
	for _, row := range roleRows {
		roleIds = append(roleIds, row.Id)
		roles = append(roles, &types.Role{Id: int(row.Id), Name: row.Name})
	}

	// 获取用户权限
	permRows, _, err := model.NewPermissionModel(l.svcCtx.Db).FindByRoleIds(l.ctx, roleIds)
	if err != nil {
		return nil, err
	}

	// 组装权限信息
	permissions := make([]*types.Permission, 0)
	permMap := map[int64]*types.Permission{}
	for _, row := range permRows {
		permMap[row.Id] = &types.Permission{Id: int(row.Id), Name: row.Name, Key: row.Key, Children: []*types.Permission{}}
	}
	for _, row := range permRows {
		if permMap[row.Pid] != nil {
			permMap[row.Pid].Children = append(permMap[row.Pid].Children, permMap[row.Id])
		}

		if row.Pid == 0 {
			permissions = append(permissions, permMap[row.Id])
		}
	}

	// 初始化用户数据
	resp.Id = int(user.Id)
	resp.Account = user.Account
	resp.Nickname = user.Nickname
	resp.Email = user.Email
	resp.Avatar = user.Avatar
	resp.LastLoginIp = user.LastLoginIp
	resp.LastLoginTime = int(utils.NowUnixSecond())
	resp.CreateTime = int(user.CreateTime)
	resp.Roles = roles
	resp.Permissions = permissions

	// 生成授权令牌
	resp.Token, err = l.getToken(resp)
	if err != nil {
		return nil, err
	}

	// 记录登陆时间
	defer func() {
		user.LastLoginTime = int64(resp.LastLoginTime)
		err = userModel.Update(l.ctx, user)
		if err != nil {
			color.Red(err.Error())
		}
	}()

	return resp, nil
}

func (l *LoginLogic) getToken(resp *types.LoginRes) (string, error) {
	marshal, err := json.Marshal(types.LoginUser{
		Id:       resp.User.Id,
		Account:  resp.User.Account,
		Nickname: resp.User.Nickname,
		IsRoot:   resp.User.IsRoot,
	})
	if err != nil {
		return "", err
	}
	tokenStr, err := token.Generate(l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessTTL, string(marshal))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
