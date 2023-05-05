package objecthashmodel

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ObjectHashModel = (*customObjectHashModel)(nil)

type (
	// ObjectHashModel is an interface to be customized, add more methods here,
	// and implement the added methods in customObjectHashModel.
	ObjectHashModel interface {
		objectHashModel
		DelBucketAllHash(ctx context.Context, bucketName string) error
	}

	customObjectHashModel struct {
		*defaultObjectHashModel
	}
)

// DelBucketAllHash 删除对应桶全部哈希
func (m *defaultObjectHashModel) DelBucketAllHash(ctx context.Context, bucketName string) error {
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		query := fmt.Sprintf("delete from %s where `hashcode` like '%s/%%'", m.table, bucketName)

		return conn.ExecCtx(ctx, query)
	})

	return err
}

// NewObjectHashModel returns a model for the database table.
func NewObjectHashModel(conn sqlx.SqlConn, c cache.CacheConf) ObjectHashModel {
	return &customObjectHashModel{
		defaultObjectHashModel: newObjectHashModel(conn, c),
	}
}
