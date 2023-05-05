package logic

import (
	"context"
	"github.com/pkg/errors"
	"oos-system/common/xerr"

	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/bucket/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemovePolicyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemovePolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemovePolicyLogic {
	return &RemovePolicyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemovePolicyLogic) RemovePolicy(in *bucket.PolicyReq) (*bucket.SucResp, error) {
	removePolicy, err := l.svcCtx.Casbin.RemovePolicy(in.Username, in.BucketName, in.Permission)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "RemovePolicy db err , username : %s , bucketName: %s err : %+v", in.Username, in.BucketName, err)
	}

	if !removePolicy {
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to RemovePolicy"), "RemovePolicy false , username : %s , bucketName: %s err : %+v", in.Username, in.BucketName, err)
	}

	return &bucket.SucResp{
		Success: removePolicy,
	}, nil
}
