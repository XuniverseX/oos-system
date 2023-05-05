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

type UpdateBucketPermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBucketPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBucketPermissionLogic {
	return &UpdateBucketPermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBucketPermissionLogic) UpdateBucketPermission(req *types.UpdateBucketPermissionReq) (resp *types.UpdateBucketPermissionResp, err error) {

	username := ctxdata.GetUserNameFromCtx(l.ctx)

	policyResp, err := l.svcCtx.BucketRpc.GetPolicy(l.ctx, &bucket.GetPolicyReq{
		Username:   username,
		BucketName: req.BucketName,
	})
	if err != nil {
		return nil, err
	}

	if policyResp.UserPermission != 3 {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(xerr.No_Bucket_Permission, "没有修改桶配置权限"), "没有修改桶配置权限 userPossess : %s, bucket: %s", username, req.BucketName)
	}

	_, err = l.svcCtx.BucketRpc.UpdateBucketPermission(l.ctx, &bucket.UpdateBucketPermission{
		BucketName: req.BucketName,
		Permission: req.Permission,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateBucketPermissionResp{}, nil
}
