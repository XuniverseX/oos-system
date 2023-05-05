package logic

import (
	"context"
	"github.com/pkg/errors"
	"oos-system/common/xerr"

	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/bucket/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelBucketAllPolicyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelBucketAllPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelBucketAllPolicyLogic {
	return &DelBucketAllPolicyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelBucketAllPolicyLogic) DelBucketAllPolicy(in *bucket.DelBucketAllPolicyReq) (*bucket.SucResp, error) {
	AllBucketPolicyList := l.svcCtx.Casbin.GetFilteredPolicy(1, in.BucketName)

	removePolicies, err := l.svcCtx.Casbin.RemovePolicies(AllBucketPolicyList)

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "remove bucket all policies db err : %v", err)
	}

	return &bucket.SucResp{
		Success: removePolicies,
	}, nil
}
