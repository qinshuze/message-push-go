package model

import (
	"ccps.com/internal/utils/arrays"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PermissionModel = (*customPermissionModel)(nil)

type (
	// PermissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPermissionModel.
	PermissionModel interface {
		permissionModel
		FindByRoleIds(ctx context.Context, roleIds []int64) (perms []*Permission, relations []*RolePermission, err error)
	}

	customPermissionModel struct {
		*defaultPermissionModel
	}
)

// NewPermissionModel returns a model for the database table.
func NewPermissionModel(conn sqlx.SqlConn) PermissionModel {
	return &customPermissionModel{
		defaultPermissionModel: newPermissionModel(conn),
	}
}

func (m *defaultPermissionModel) FindByRoleIds(ctx context.Context, roleIds []int64) (perms []*Permission, relations []*RolePermission, err error) {
	var ids []int64

	relations, err = NewRolePermissionModel(m.conn).FindByRoleIds(ctx, roleIds)
	if err != nil {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}

	if len(relations) == 0 {
		return perms, relations, nil
	}

	for _, rel := range relations {
		ids = append(ids, rel.PermissionId)
	}

	var queryStr = fmt.Sprintf("select %s from %s where id in(%s) ", permissionRows, m.table, arrays.Int64ToString(ids))
	err = m.conn.QueryRowsCtx(ctx, &perms, queryStr)
	if err != nil {
		return nil, nil, err
	}

	return perms, relations, nil
}
