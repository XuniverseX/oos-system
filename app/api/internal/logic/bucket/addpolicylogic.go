package bucket

import (
	"context"
	"github.com/pkg/errors"
	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/usercenter/usercenter"
	"oos-system/common/ctxdata"
	"oos-system/common/xerr"

	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPolicyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPolicyLogic {
	return &AddPolicyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPolicyLogic) AddPolicy(req *types.AddPolicyReq) (resp *types.AddPolicyResp, err error) {
	username := ctxdata.GetUserNameFromCtx(l.ctx)

	policyResp, getPolicyErr := l.svcCtx.BucketRpc.GetPolicy(l.ctx, &bucket.GetPolicyReq{
		Username:   username,
		BucketName: req.BucketName,
	})
	if getPolicyErr != nil {
		return nil, getPolicyErr
	}

	if policyResp.UserPermission != 3 {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.No_Bucket_Permission), "添加桶协作者没有权限")
	}

	// 查询用户是否存在
	_, isExitUserErr := l.svcCtx.UsercenterRpc.IsExitUser(l.ctx, &usercenter.IsExitUserReq{
		Username: req.UserName,
	})

	if isExitUserErr != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户不存在"), "添加桶权限 用户不存在 %v", isExitUserErr)
	}

	_, addPolicyErr := l.svcCtx.BucketRpc.AddPolicy(l.ctx, &bucket.PolicyReq{
		Username:   req.UserName,
		BucketName: req.BucketName,
		Permission: req.UserPermission,
	})
	if addPolicyErr != nil {
		return nil, addPolicyErr
	}

	return &types.AddPolicyResp{}, nil
}
