package logic

import (
	"context"
	"github.com/pkg/errors"
	"math"
	"oos-system/app/rpc/model/bucketmodel"
	"oos-system/common/xerr"
	"strconv"
	"strings"

	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/bucket/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyBucketByBucketNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMyBucketByBucketNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyBucketByBucketNameLogic {
	return &GetMyBucketByBucketNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMyBucketByBucketNameLogic) GetMyBucketByBucketName(in *bucket.GetBucketByBucketNameReq) (*bucket.AllBucketResp, error) {
	// 查询casbin获取用户所有拥有操作权限的桶

	//获取用户所有权限策略
	policyList := l.svcCtx.Casbin.GetFilteredPolicy(0, in.Username)
	// 获取自己的桶
	var myBucketList []string

	for _, item := range policyList {
		if item[2] == "3" {
			myBucketList = append(myBucketList, item[1])
		}
	}

	// kmp子串模糊匹配搜索的桶
	var likeBucketList []string
	for _, item := range myBucketList {
		if strings.Contains(item, in.BucketName) {
			likeBucketList = append(likeBucketList, item)
		}
	}

	// 根据桶名查询桶
	whereBuilder := l.svcCtx.BucketModel.RowBuilder()

	list, err := l.svcCtx.BucketModel.FindBucketList(l.ctx, whereBuilder, likeBucketList)
	if err != nil && err != bucketmodel.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to get user's bucket list err : %v , in :%+v", err, in)
	}

	var resp []*bucket.BucketInfoResp

	if len(list) > 0 {
		for _, bucketItem := range list {
			var bucketInfo bucket.BucketInfoResp

			bucketInfo.Id = bucketItem.Id
			bucketInfo.Username = bucketItem.Username
			bucketInfo.BucketName = bucketItem.BucketName
			// 保留整数 截断小数
			value := int(math.Floor(bucketItem.CapacityCur))
			bucketInfo.CapacityCur = strconv.Itoa(value)
			bucketInfo.ObjectNum = bucketItem.ObjectNum
			bucketInfo.Permission = bucketItem.Permission
			bucketInfo.UserPermission = 3 // 3 代表桶创建则 1 读写 2 读 3 写
			bucketInfo.CreateTime = bucketItem.CreateTime.Format("2006-01-02 15:04:05")
			bucketInfo.UpdateTime = bucketItem.UpdateTime.Format("2006-01-02 15:04:05")

			resp = append(resp, &bucketInfo)
		}
	}

	return &bucket.AllBucketResp{
		BucketList: resp,
	}, nil
}
