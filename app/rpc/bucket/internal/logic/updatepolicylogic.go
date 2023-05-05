package logic

import (
	"context"
	"github.com/pkg/errors"
	"oos-system/common/xerr"

	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/bucket/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePolicyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePolicyLogic {
	return &UpdatePolicyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePolicyLogic) UpdatePolicy(in *bucket.UpdatePolicyReq) (*bucket.SucResp, error) {
	updatePolicy, err := l.svcCtx.Casbin.UpdatePolicy([]string{in.Username, in.BucketName, in.OldPermission}, []string{in.Username, in.BucketName, in.NewPermission})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "UpdatePolicy db err , username : %s , bucketName: %s err : %+v", in.Username, in.BucketName, err)
	}

	if !updatePolicy {
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to UpdatePolicy"), "UpdatePolicy false , username : %s , bucketName: %s err : %+v", in.Username, in.BucketName, err)
	}
	return &bucket.SucResp{
		Success: updatePolicy,
	}, nil
}
