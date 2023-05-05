package logic

import (
	"context"

	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/bucket/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllBucketPolicyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllBucketPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllBucketPolicyLogic {
	return &GetAllBucketPolicyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllBucketPolicyLogic) GetAllBucketPolicy(in *bucket.GetAllBucketPolicyReq) (*bucket.GetAllBucketPolicyResp, error) {
	policyList := l.svcCtx.Casbin.GetFilteredPolicy(1, in.BucketName)

	var resp []*bucket.UserPolicyInfo

	if len(policyList) > 0 {
		for _, policyItem := range policyList {
			var userPolicy bucket.UserPolicyInfo
			userPolicy.Username = policyItem[0]
			userPolicy.BucketName = policyItem[1]
			userPolicy.UserPermission = policyItem[2]

			resp = append(resp, &userPolicy)
		}
	}

	return &bucket.GetAllBucketPolicyResp{
		PolicyList: resp,
	}, nil
}
