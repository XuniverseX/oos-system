// Code generated by goctl. DO NOT EDIT.

package objecthashmodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	objectHashFieldNames          = builder.RawFieldNames(&ObjectHash{})
	objectHashRows                = strings.Join(objectHashFieldNames, ",")
	objectHashRowsExpectAutoSet   = strings.Join(stringx.Remove(objectHashFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	objectHashRowsWithPlaceHolder = strings.Join(stringx.Remove(objectHashFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheObjectHashIdPrefix       = "cache:objectHash:id:"
	cacheObjectHashHashcodePrefix = "cache:objectHash:hashcode:"
)

type (
	objectHashModel interface {
		Insert(ctx context.Context, data *ObjectHash) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ObjectHash, error)
		FindOneByHashcode(ctx context.Context, hashcode string) (*ObjectHash, error)
		Update(ctx context.Context, data *ObjectHash) error
		Delete(ctx context.Context, id int64) error
	}

	defaultObjectHashModel struct {
		sqlc.CachedConn
		table string
	}

	ObjectHash struct {
		Id       int64  `db:"id"`       // id
		Hashcode string `db:"hashcode"` // 对象hashcode
	}
)

func newObjectHashModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultObjectHashModel {
	return &defaultObjectHashModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`object_hash`",
	}
}

func (m *defaultObjectHashModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	objectHashHashcodeKey := fmt.Sprintf("%s%v", cacheObjectHashHashcodePrefix, data.Hashcode)
	objectHashIdKey := fmt.Sprintf("%s%v", cacheObjectHashIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, objectHashHashcodeKey, objectHashIdKey)
	return err
}

func (m *defaultObjectHashModel) FindOne(ctx context.Context, id int64) (*ObjectHash, error) {
	objectHashIdKey := fmt.Sprintf("%s%v", cacheObjectHashIdPrefix, id)
	var resp ObjectHash
	err := m.QueryRowCtx(ctx, &resp, objectHashIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", objectHashRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultObjectHashModel) FindOneByHashcode(ctx context.Context, hashcode string) (*ObjectHash, error) {
	objectHashHashcodeKey := fmt.Sprintf("%s%v", cacheObjectHashHashcodePrefix, hashcode)
	var resp ObjectHash
	err := m.QueryRowIndexCtx(ctx, &resp, objectHashHashcodeKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `hashcode` = ? limit 1", objectHashRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, hashcode); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultObjectHashModel) Insert(ctx context.Context, data *ObjectHash) (sql.Result, error) {
	objectHashHashcodeKey := fmt.Sprintf("%s%v", cacheObjectHashHashcodePrefix, data.Hashcode)
	objectHashIdKey := fmt.Sprintf("%s%v", cacheObjectHashIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, objectHashRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.Hashcode)
	}, objectHashHashcodeKey, objectHashIdKey)
	return ret, err
}

func (m *defaultObjectHashModel) Update(ctx context.Context, newData *ObjectHash) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	objectHashHashcodeKey := fmt.Sprintf("%s%v", cacheObjectHashHashcodePrefix, data.Hashcode)
	objectHashIdKey := fmt.Sprintf("%s%v", cacheObjectHashIdPrefix, data.Id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, objectHashRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Hashcode, newData.Id)
	}, objectHashHashcodeKey, objectHashIdKey)
	return err
}

func (m *defaultObjectHashModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheObjectHashIdPrefix, primary)
}

func (m *defaultObjectHashModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", objectHashRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultObjectHashModel) tableName() string {
	return m.table
}
