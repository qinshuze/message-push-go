// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	rolePermissionFieldNames          = builder.RawFieldNames(&RolePermission{}, true)
	rolePermissionRows                = strings.Join(rolePermissionFieldNames, ",")
	rolePermissionRowsExpectAutoSet   = strings.Join(stringx.Remove(rolePermissionFieldNames, "id"), ",")
	rolePermissionRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(rolePermissionFieldNames, "id"))
)

type (
	rolePermissionModel interface {
		Insert(ctx context.Context, data *RolePermission) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*RolePermission, error)
		Update(ctx context.Context, data *RolePermission) error
		Delete(ctx context.Context, id int64) error
	}

	defaultRolePermissionModel struct {
		conn  sqlx.SqlConn
		table string
	}

	RolePermission struct {
		Id           int64 `db:"id"`
		RoleId       int64 `db:"role_id"`       // 角色id
		PermissionId int64 `db:"permission_id"` // 权限id
		Scope        int64 `db:"scope"`         // 权限适用范围：1 - 全局，2 - 仅适用于登陆用户自己的数据范围内
	}
)

func newRolePermissionModel(conn sqlx.SqlConn) *defaultRolePermissionModel {
	return &defaultRolePermissionModel{
		conn:  conn,
		table: `"public"."role_permission"`,
	}
}

func (m *defaultRolePermissionModel) withSession(session sqlx.Session) *defaultRolePermissionModel {
	return &defaultRolePermissionModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: `"public"."role_permission"`,
	}
}

func (m *defaultRolePermissionModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where id = $1", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultRolePermissionModel) FindOne(ctx context.Context, id int64) (*RolePermission, error) {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", rolePermissionRows, m.table)
	var resp RolePermission
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRolePermissionModel) Insert(ctx context.Context, data *RolePermission) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3)", m.table, rolePermissionRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.RoleId, data.PermissionId, data.Scope)
	return ret, err
}

func (m *defaultRolePermissionModel) Update(ctx context.Context, data *RolePermission) error {
	query := fmt.Sprintf("update %s set %s where id = $1", m.table, rolePermissionRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Id, data.RoleId, data.PermissionId, data.Scope)
	return err
}

func (m *defaultRolePermissionModel) tableName() string {
	return m.table
}