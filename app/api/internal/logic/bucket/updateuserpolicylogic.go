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

type UpdateUserPolicyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserPolicyLogic {
	return &UpdateUserPolicyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserPolicyLogic) UpdateUserPolicy(req *types.UpdatePolicyReq) (resp *types.UpdatePolicyResp, err error) {
	// 查询是不是桶的拥有者 桶的拥有者才可以修改协作者权限
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

	// 修改协作者权限
	_, updatePolicyErr := l.svcCtx.BucketRpc.UpdatePolicy(l.ctx, &bucket.UpdatePolicyReq{
		Username:      req.UserName,
		BucketName:    req.BucketName,
		OldPermission: req.OldPermission,
		NewPermission: req.NewPermission,
	})

	if updatePolicyErr != nil {
		return nil, updatePolicyErr
	}

	return &types.UpdatePolicyResp{}, nil
}
