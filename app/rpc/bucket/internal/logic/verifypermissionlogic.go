package logic

import (
	"context"
	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/bucket/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyPermissionLogic {
	return &VerifyPermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyPermissionLogic) VerifyPermission(in *bucket.PolicyReq) (*bucket.SucResp, error) {
	hasPolicy := l.svcCtx.Casbin.HasPolicy(in.Username, in.BucketName, in.Permission)

	//if !hasPolicy {
	//	return nil, errors.Wrapf(xerr.NewErrMsg("no exist this policy"), "no exist this policy , username : %s , bucketName: %s", in.Username, in.BucketName)
	//}

	return &bucket.SucResp{
		Success: hasPolicy,
	}, nil
}
