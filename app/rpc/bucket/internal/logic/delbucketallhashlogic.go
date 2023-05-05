package logic

import (
	"context"

	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/bucket/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelBucketAllHashLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelBucketAllHashLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelBucketAllHashLogic {
	return &DelBucketAllHashLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelBucketAllHashLogic) DelBucketAllHash(in *bucket.DelBucketAllHashReq) (*bucket.SucResp, error) {
	// 删除桶里所有hash
	err := l.svcCtx.ObjectHashMode.DelBucketAllHash(l.ctx, in.BucketName)

	if err != nil {
		return nil, err
	}

	return &bucket.SucResp{}, nil
}
