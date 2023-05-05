package bucket

import (
	"context"
	"oos-system/app/rpc/bucket/bucket"
	"oos-system/common/ctxdata"

	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyBucketByBucketNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMyBucketByBucketNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyBucketByBucketNameLogic {
	return &GetMyBucketByBucketNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyBucketByBucketNameLogic) GetMyBucketByBucketName(req *types.GetBucketByBucketNameReq) (resp *types.GetAllBucketResq, err error) {
	username := ctxdata.GetUserNameFromCtx(l.ctx)

	// 调用rpc
	likeBucketList, err := l.svcCtx.BucketRpc.GetMyBucketByBucketName(l.ctx, &bucket.GetBucketByBucketNameReq{
		Username:   username,
		BucketName: req.BucketName,
	})
	if err != nil {
		return nil, err
	}

	var typesGetAllBucketByUserNameList []types.BucketInfo

	if len(likeBucketList.BucketList) > 0 {
		for _, bucketItem := range likeBucketList.BucketList {
			var typesGetAllBucketByUserName types.BucketInfo

			typesGetAllBucketByUserName.Id = bucketItem.Id
			typesGetAllBucketByUserName.Username = bucketItem.Username
			typesGetAllBucketByUserName.BucketName = bucketItem.BucketName
			typesGetAllBucketByUserName.CapacityCur = bucketItem.CapacityCur
			typesGetAllBucketByUserName.ObjectNum = bucketItem.ObjectNum
			typesGetAllBucketByUserName.Premission = bucketItem.Permission
			typesGetAllBucketByUserName.UserPremission = bucketItem.UserPermission
			typesGetAllBucketByUserName.CreateTime = bucketItem.CreateTime
			typesGetAllBucketByUserName.UpdateTime = bucketItem.UpdateTime

			typesGetAllBucketByUserNameList = append(typesGetAllBucketByUserNameList, typesGetAllBucketByUserName)
		}
	}

	return &types.GetAllBucketResq{
		List: typesGetAllBucketByUserNameList,
	}, nil
}
