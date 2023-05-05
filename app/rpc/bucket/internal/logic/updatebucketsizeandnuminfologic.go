package logic

import (
	"context"
	"github.com/pkg/errors"
	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/bucket/internal/svc"
	"oos-system/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBucketSizeAndNumInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBucketSizeAndNumInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBucketSizeAndNumInfoLogic {
	return &UpdateBucketSizeAndNumInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateBucketSizeAndNumInfoLogic) UpdateBucketSizeAndNumInfo(in *bucket.UpdateBucketSizeAndNumReq) (*bucket.SucResp, error) {
	//bucketItem, err := l.svcCtx.BucketModel.FindOneByBucketName(l.ctx, in.BucketName)
	//
	//if err != nil {
	//	return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db 根据桶名获取桶失败 err: %v", err)
	//}

	//bucketItem.CapacityCur += float64(in.Size) / 1024
	//// 这个有精度问题 可能会减到小于0 判断一下
	//if bucketItem.CapacityCur < 0 {
	//	bucketItem.CapacityCur = 0.0
	//}
	//bucketItem.ObjectNum += in.ObjectNum
	//fmt.Println("bucketNum: ", bucketItem.ObjectNum)
	//bucketItem.UpdateTime = time.Now()

	//insertErr := l.svcCtx.BucketModel.Update(l.ctx, bucketItem)

	//if insertErr != nil {
	//	return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db 更新桶信息失败 err: %v", insertErr)
	//}
	size := float64(in.Size) / 1024

	err := l.svcCtx.BucketModel.UpdateSizeAndNum(l.ctx, size, in.ObjectNum, in.BucketName)

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db 更新桶信息失败 err: %v", err)
	}

	return &bucket.SucResp{
		Success: true,
	}, nil
}
