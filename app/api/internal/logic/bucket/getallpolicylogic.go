package bucket

import (
	"context"
	"oos-system/app/rpc/bucket/bucket"

	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllPolicyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllPolicyLogic {
	return &GetAllPolicyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllPolicyLogic) GetAllPolicy(req *types.GetAllPolicyReq) (resp *types.GetAllPolicyResp, err error) {
	allBucketPolicy, err := l.svcCtx.BucketRpc.GetAllBucketPolicy(l.ctx, &bucket.GetAllBucketPolicyReq{
		BucketName: req.BucketName,
	})
	if err != nil {
		return nil, err
	}

	var allBucketPolicyList []types.PolicyInfo

	if len(allBucketPolicy.PolicyList) > 0 {
		for _, policyItem := range allBucketPolicy.PolicyList {
			var policyInfo types.PolicyInfo

			policyInfo.UserName = policyItem.Username
			policyInfo.BucketName = policyItem.BucketName
			policyInfo.UserPermission = policyItem.UserPermission

			allBucketPolicyList = append(allBucketPolicyList, policyInfo)
		}
	}

	return &types.GetAllPolicyResp{
		List: allBucketPolicyList,
	}, nil
}
