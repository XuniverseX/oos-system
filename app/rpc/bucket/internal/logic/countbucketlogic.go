package logic

import (
	"context"
	"github.com/pkg/errors"
	"oos-system/app/rpc/model/bucketmodel"
	"oos-system/common/xerr"

	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/bucket/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountBucketLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCountBucketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountBucketLogic {
	return &CountBucketLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CountBucketLogic) CountBucket(in *bucket.CountBucketReq) (*bucket.CountBucketResp, error) {

	// 生成count(bucket_name)字段 squirrel
	countBuilder := l.svcCtx.BucketModel.CountBuilder("bucket_name")

	// 调用sql语句
	bucketNum, err := l.svcCtx.BucketModel.CountBucket(l.ctx, countBuilder, in.Username)

	if err != nil && err != bucketmodel.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to count user's bucket err : %v , username :%+v", err, in.Username)
	}

	return &bucket.CountBucketResp{
		BucketNum: bucketNum,
	}, nil
}
