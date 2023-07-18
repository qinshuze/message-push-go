package model

import (
	"ccps.com/internal/utils/arrays"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RolePermissionModel = (*customRolePermissionModel)(nil)

type (
	// RolePermissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRolePermissionModel.
	RolePermissionModel interface {
		rolePermissionModel
		FindByRoleIds(ctx context.Context, roleIds []int64) ([]*RolePermission, error)
	}

	customRolePermissionModel struct {
		*defaultRolePermissionModel
	}
)

// NewRolePermissionModel returns a model for the database table.
func NewRolePermissionModel(conn sqlx.SqlConn) RolePermissionModel {
	return &customRolePermissionModel{
		defaultRolePermissionModel: newRolePermissionModel(conn),
	}
}

func (m *defaultRolePermissionModel) FindByRoleIds(ctx context.Context, roleIds []int64) ([]*RolePermission, error) {
	var rolePermissions []*RolePermission
	var queryStr = fmt.Sprintf("select %s from %s where role_id in(%s) ", rolePermissionRows, m.table, arrays.Int64ToString(roleIds))

	err := m.conn.QueryRowsCtx(ctx, &rolePermissions, queryStr)
	if err != nil {
		return nil, err
	}

	return rolePermissions, nil
}
