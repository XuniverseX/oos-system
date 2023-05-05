package bucket

import (
	"context"
	"github.com/pkg/errors"
	"oos-system/app/rpc/bucket/bucket"
	"oos-system/common/ctxdata"
	"oos-system/common/xerr"

	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveUserPolicyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveUserPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveUserPolicyLogic {
	return &RemoveUserPolicyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveUserPolicyLogic) RemoveUserPolicy(req *types.RemovePolicyReq) (resp *types.RemovePolicyResp, err error) {
	// 查询是不是桶的拥有者 桶的拥有者才可以删除协作者权限
	userPossess := ctxdata.GetUserNameFromCtx(l.ctx)

	possessPolicy, policyErr := l.svcCtx.BucketRpc.GetPolicy(l.ctx, &bucket.GetPolicyReq{
		Username:   userPossess,
		BucketName: req.BucketName,
	})
	if policyErr != nil {
		return nil, policyErr
	}
	if possessPolicy.UserPermission != 3 {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(xerr.No_Bucket_Permission, "没有修改协作者管理权权限"), "没有修改协作者管理权权限 userPossess : %s, bucket: %s", userPossess, req.BucketName)
	}

	// 删除权限
	_, removerPolicyErr := l.svcCtx.BucketRpc.RemovePolicy(l.ctx, &bucket.PolicyReq{
		Username:   req.UserName,
		BucketName: req.BucketName,
		Permission: req.UserPermission,
	})

	if removerPolicyErr != nil {
		return nil, removerPolicyErr
	}

	return &types.RemovePolicyResp{}, nil
}
