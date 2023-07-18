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
	permissionFieldNames          = builder.RawFieldNames(&Permission{}, true)
	permissionRows                = strings.Join(permissionFieldNames, ",")
	permissionRowsExpectAutoSet   = strings.Join(stringx.Remove(permissionFieldNames, "id"), ",")
	permissionRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(permissionFieldNames, "id"))
)

type (
	permissionModel interface {
		Insert(ctx context.Context, data *Permission) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Permission, error)
		FindOneByKey(ctx context.Context, key string) (*Permission, error)
		Update(ctx context.Context, data *Permission) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPermissionModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Permission struct {
		Id         int64  `db:"id"`
		Name       string `db:"name"`        // 名称
		Key        string `db:"key"`         // 键
		Route      string `db:"route"`       // 资源路由
		Remark     string `db:"remark"`      // 备注
		CreateTime int64  `db:"create_time"` // 创建时间
		CreatorId  int64  `db:"creator_id"`  // 创建人id
		Pid        int64  `db:"pid"`         // 父id
		Realm      string `db:"realm"`       // 所属域
	}
)

func newPermissionModel(conn sqlx.SqlConn) *defaultPermissionModel {
	return &defaultPermissionModel{
		conn:  conn,
		table: `"public"."permission"`,
	}
}

func (m *defaultPermissionModel) withSession(session sqlx.Session) *defaultPermissionModel {
	return &defaultPermissionModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: `"public"."permission"`,
	}
}

func (m *defaultPermissionModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where id = $1", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPermissionModel) FindOne(ctx context.Context, id int64) (*Permission, error) {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", permissionRows, m.table)
	var resp Permission
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

func (m *defaultPermissionModel) FindOneByKey(ctx context.Context, key string) (*Permission, error) {
	var resp Permission
	query := fmt.Sprintf("select %s from %s where key = $1 limit 1", permissionRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, key)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPermissionModel) Insert(ctx context.Context, data *Permission) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7, $8)", m.table, permissionRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.Key, data.Route, data.Remark, data.CreateTime, data.CreatorId, data.Pid, data.Realm)
	return ret, err
}

func (m *defaultPermissionModel) Update(ctx context.Context, newData *Permission) error {
	query := fmt.Sprintf("update %s set %s where id = $1", m.table, permissionRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Id, newData.Name, newData.Key, newData.Route, newData.Remark, newData.CreateTime, newData.CreatorId, newData.Pid, newData.Realm)
	return err
}

func (m *defaultPermissionModel) tableName() string {
	return m.table
}