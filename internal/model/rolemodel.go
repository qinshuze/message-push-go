package model

import (
	"ccps.com/internal/utils/arrays"
	"context"
	"fmt"
	zsqlx "github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RoleModel = (*customRoleModel)(nil)

type (
	// RoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoleModel.
	RoleModel interface {
		roleModel
		FindByUserIds(ctx context.Context, userIds []int64) (roles []*Role, relations []*RoleUser, err error)
	}

	customRoleModel struct {
		*defaultRoleModel
	}
)

// NewRoleModel returns a model for the database table.
func NewRoleModel(conn zsqlx.SqlConn) RoleModel {
	return &customRoleModel{
		defaultRoleModel: newRoleModel(conn),
	}
}

func (m *defaultRoleModel) FindByUserIds(ctx context.Context, userIds []int64) (roles []*Role, relations []*RoleUser, err error) {
	roles = make([]*Role, 0)

	relations, err = NewRoleUserModel(m.conn).FindByUserIds(ctx, userIds)
	if err != nil {
		return nil, nil, err
	}

	if len(relations) == 0 {
		return roles, relations, nil
	}

	var ids []int64
	for _, rel := range relations {
		ids = append(ids, rel.RoleId)
	}

	queryStr := fmt.Sprintf("select %s from %s where id in(%s)", roleRows, m.table, arrays.Int64ToString(ids))
	err = m.conn.QueryRowsCtx(ctx, &roles, queryStr)
	if err != nil {
		return nil, nil, err
	}

	return roles, relations, nil
}
