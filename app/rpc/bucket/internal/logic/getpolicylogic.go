package logic

import (
	"context"
	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/bucket/internal/svc"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPolicyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPolicyLogic {
	return &GetPolicyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPolicyLogic) GetPolicy(in *bucket.GetPolicyReq) (*bucket.GetPolicyResp, error) {
	filteredPolicy := l.svcCtx.Casbin.GetFilteredPolicy(1, in.BucketName)

	var userPermission int64 = -1

	if len(filteredPolicy) != 0 {
		for index, policyItem := range filteredPolicy {
			if policyItem[0] == in.Username {
				userPermission, _ = strconv.ParseInt(filteredPolicy[index][2], 10, 64)
			}
		}
	}

	return &bucket.GetPolicyResp{
		UserPermission: userPermission,
	}, nil
}
