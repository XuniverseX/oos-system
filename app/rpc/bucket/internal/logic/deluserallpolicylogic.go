package logic

import (
	"context"
	"github.com/pkg/errors"
	"oos-system/common/xerr"

	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/bucket/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserAllPolicyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserAllPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserAllPolicyLogic {
	return &DelUserAllPolicyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserAllPolicyLogic) DelUserAllPolicy(in *bucket.DelUserAllPolicy) (*bucket.SucResp, error) {
	AllBucketPolicyList := l.svcCtx.Casbin.GetFilteredPolicy(0, in.UserName)

	removeUserPolicies, err := l.svcCtx.Casbin.RemovePolicies(AllBucketPolicyList)

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "remove user All policies db err : %v", err)
	}

	return &bucket.SucResp{
		Success: removeUserPolicies,
	}, nil
}
