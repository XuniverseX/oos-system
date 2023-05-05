package logic

import (
	"context"
	"github.com/pkg/errors"
	"oos-system/common/xerr"

	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/bucket/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPolicyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPolicyLogic {
	return &AddPolicyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddPolicyLogic) AddPolicy(in *bucket.PolicyReq) (*bucket.SucResp, error) {

	// 在添加策略之前先查询该桶下已经添加的权限人数 不得超过五个人
	exitsPolices := l.svcCtx.Casbin.GetFilteredPolicy(1, in.BucketName)

	if len(exitsPolices) > 5 {
		return nil, errors.Wrapf(xerr.NewErrMsg("权限添加失败，协作管理桶成员超过五人"), "%s桶协作超过五人", in.BucketName)
	}

	// 添加权限
	// 0 读写权限 1 读权限 2 写权限
	ok, err := l.svcCtx.Casbin.AddPolicy(in.Username, in.BucketName, in.Permission)
	if err != nil {
		return nil, errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "权限添加失败")
	}

	return &bucket.SucResp{
		Success: ok,
	}, nil
}
