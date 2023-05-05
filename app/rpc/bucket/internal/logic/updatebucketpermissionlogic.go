package logic

import (
	"context"
	"strconv"

	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/bucket/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBucketPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBucketPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBucketPermissionLogic {
	return &UpdateBucketPermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateBucketPermissionLogic) UpdateBucketPermission(in *bucket.UpdateBucketPermission) (*bucket.SucResp, error) {
	bucketItem, err := l.svcCtx.BucketModel.FindOneByBucketName(l.ctx, in.BucketName)
	if err != nil {
		return nil, err
	}
	bucketItem.Permission, _ = strconv.ParseInt(in.Permission, 10, 64)

	_ = l.svcCtx.BucketModel.Update(l.ctx, bucketItem)

	return &bucket.SucResp{}, nil
}
