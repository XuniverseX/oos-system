package bucketmodel

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BucketModel = (*customBucketModel)(nil)

type (
	// BucketModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBucketModel.
	BucketModel interface {
		bucketModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		RowBuilder() squirrel.SelectBuilder
		DeleteByBucketName(ctx context.Context, bucketName string) error
		CountBuilder(filed string) squirrel.SelectBuilder
		FindBucketList(ctx context.Context, rowBuilder squirrel.SelectBuilder, bucketNames []string) ([]*Bucket, error)
		CountBucket(ctx context.Context, countBuilder squirrel.SelectBuilder, username string) (int64, error)
		UpdateSizeAndNum(ctx context.Context, size float64, num int64, bucketName string) error
	}

	customBucketModel struct {
		*defaultBucketModel
	}
)

// NewBucketModel returns a model for the database table.
func NewBucketModel(conn sqlx.SqlConn, c cache.CacheConf) BucketModel {
	return &customBucketModel{
		defaultBucketModel: newBucketModel(conn, c),
	}
}

func (m *defaultBucketModel) DeleteByBucketName(ctx context.Context, bucketName string) error {

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		query := fmt.Sprintf("delete from %s where `bucket_name` = ?", m.table)

		return conn.ExecCtx(ctx, query, bucketName)
	})

	return err
}

func (m *defaultBucketModel) CountBucket(ctx context.Context, countBuilder squirrel.SelectBuilder, username string) (int64, error) {
	query, values, err := countBuilder.Where("username = ?", username).ToSql()
	if err != nil {
		// 查询错误返回-1
		return -1, err
	}

	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return -1, err
	}
}

func (m *defaultBucketModel) FindBucketList(ctx context.Context, rowBuilder squirrel.SelectBuilder, bucketNames []string) ([]*Bucket, error) {
	rowBuilder = rowBuilder.OrderBy("id DESC")
	query, values, err := rowBuilder.Where(squirrel.Eq{"bucket_name": bucketNames}).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Bucket
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultBucketModel) UpdateSizeAndNum(ctx context.Context, size float64, num int64, bucketName string) error {
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		// 我好害怕他会嘎蛋
		query := fmt.Sprintf("update %s set `capacity_cur` = IF(`capacity_cur` + %f < 0, 0.0, `capacity_cur` + %f),`object_num` = `object_num` + %d, `update_time` = NOW() where `bucket_name` = '%s'", m.table, size, size, num, bucketName)

		return conn.ExecCtx(ctx, query)
	})
	return err
}

// export logic
func (m *defaultBucketModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

func (m *defaultBucketModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultBucketModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(bucketRows).From(m.table)
}
