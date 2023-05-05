package logic

import (
	"context"
	"github.com/pkg/errors"
	"oos-system/common/xerr"

	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/bucket/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelBucketLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelBucketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelBucketLogic {
	return &DelBucketLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelBucketLogic) DelBucket(in *bucket.DelBucketReq) (*bucket.SucResp, error) {
	// 删除库里桶
	err := l.svcCtx.BucketModel.DeleteByBucketName(l.ctx, in.BucketName)

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "del bucket db err : %v", err)
	}

	return &bucket.SucResp{
		Success: err == nil,
	}, nil
}
