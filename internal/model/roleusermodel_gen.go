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
	roleUserFieldNames          = builder.RawFieldNames(&RoleUser{}, true)
	roleUserRows                = strings.Join(roleUserFieldNames, ",")
	roleUserRowsExpectAutoSet   = strings.Join(stringx.Remove(roleUserFieldNames, "id"), ",")
	roleUserRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(roleUserFieldNames, "id"))
)

type (
	roleUserModel interface {
		Insert(ctx context.Context, data *RoleUser) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*RoleUser, error)
		Update(ctx context.Context, data *RoleUser) error
		Delete(ctx context.Context, id int64) error
	}

	defaultRoleUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	RoleUser struct {
		Id     int64 `db:"id"`
		RoleId int64 `db:"role_id"` // 角色id
		UserId int64 `db:"user_id"` // 用户id
	}
)

func newRoleUserModel(conn sqlx.SqlConn) *defaultRoleUserModel {
	return &defaultRoleUserModel{
		conn:  conn,
		table: `"public"."role_user"`,
	}
}

func (m *defaultRoleUserModel) withSession(session sqlx.Session) *defaultRoleUserModel {
	return &defaultRoleUserModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: `"public"."role_user"`,
	}
}

func (m *defaultRoleUserModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where id = $1", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultRoleUserModel) FindOne(ctx context.Context, id int64) (*RoleUser, error) {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", roleUserRows, m.table)
	var resp RoleUser
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

func (m *defaultRoleUserModel) Insert(ctx context.Context, data *RoleUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values ($1, $2)", m.table, roleUserRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.RoleId, data.UserId)
	return ret, err
}

func (m *defaultRoleUserModel) Update(ctx context.Context, data *RoleUser) error {
	query := fmt.Sprintf("update %s set %s where id = $1", m.table, roleUserRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Id, data.RoleId, data.UserId)
	return err
}

func (m *defaultRoleUserModel) tableName() string {
	return m.table
}
