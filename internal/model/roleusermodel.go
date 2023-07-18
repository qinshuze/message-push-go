package model

import (
	"ccps.com/internal/utils/arrays"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RoleUserModel = (*customRoleUserModel)(nil)

type (
	// RoleUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoleUserModel.
	RoleUserModel interface {
		roleUserModel
		FindByUserIds(ctx context.Context, userIds []int64) (roleUsers []*RoleUser, err error)
	}

	customRoleUserModel struct {
		*defaultRoleUserModel
	}
)

// NewRoleUserModel returns a model for the database table.
func NewRoleUserModel(conn sqlx.SqlConn) RoleUserModel {
	return &customRoleUserModel{
		defaultRoleUserModel: newRoleUserModel(conn),
	}
}

func (m *defaultRoleUserModel) FindByUserIds(ctx context.Context, userIds []int64) (roleUsers []*RoleUser, err error) {
	roleUsers = make([]*RoleUser, 0)

	queryStr := fmt.Sprintf("select %s from %s where user_id in(%s) ", roleUserRows, m.table, arrays.Int64ToString(userIds))
	err = m.conn.QueryRowsCtx(ctx, &roleUsers, queryStr)
	if err != nil {
		return nil, err
	}

	return roleUsers, nil
}
