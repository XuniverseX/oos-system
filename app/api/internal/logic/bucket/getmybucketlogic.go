package bucket

import (
	"context"
	"oos-system/app/rpc/bucket/bucket"
	"oos-system/common/ctxdata"

	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyBucketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMyBucketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyBucketLogic {
	return &GetMyBucketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyBucketLogic) GetMyBucket(req *types.GetAllBucketReq) (resp *types.GetAllBucketResq, err error) {
	username := ctxdata.GetUserNameFromCtx(l.ctx)

	rpcRes, err := l.svcCtx.BucketRpc.GetMyBucket(l.ctx, &bucket.GetMyBucketReq{
		Username: username,
	})
	if err != nil {
		return nil, err
	}

	var typesGetAllBucketByUserNameList []types.BucketInfo

	if len(rpcRes.BucketList) > 0 {
		for _, bucket := range rpcRes.BucketList {
			var typesGetAllBucketByUserName types.BucketInfo

			typesGetAllBucketByUserName.Id = bucket.Id
			typesGetAllBucketByUserName.Username = bucket.Username
			typesGetAllBucketByUserName.BucketName = bucket.BucketName
			typesGetAllBucketByUserName.CapacityCur = bucket.CapacityCur
			typesGetAllBucketByUserName.ObjectNum = bucket.ObjectNum
			typesGetAllBucketByUserName.Premission = bucket.Permission
			typesGetAllBucketByUserName.UserPremission = bucket.UserPermission
			typesGetAllBucketByUserName.CreateTime = bucket.CreateTime
			typesGetAllBucketByUserName.UpdateTime = bucket.UpdateTime

			typesGetAllBucketByUserNameList = append(typesGetAllBucketByUserNameList, typesGetAllBucketByUserName)
		}
	}

	return &types.GetAllBucketResq{
		List: typesGetAllBucketByUserNameList,
	}, nil
}
