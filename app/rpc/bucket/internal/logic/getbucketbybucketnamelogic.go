package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"oos-system/app/rpc/model/bucketmodel"
	"oos-system/common/xerr"
	"strconv"
	"strings"

	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/bucket/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBucketByBucketNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBucketByBucketNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBucketByBucketNameLogic {
	return &GetBucketByBucketNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBucketByBucketNameLogic) GetBucketByBucketName(in *bucket.GetBucketByBucketNameReq) (*bucket.AllBucketResp, error) {
	// 查询casbin获取用户所有拥有操作权限的桶

	//获取用户拥有权限的所有桶
	policyList := l.svcCtx.Casbin.GetFilteredPolicy(0, in.Username)

	// kmp子串模糊匹配搜索的桶
	var likeBucketList []string
	for _, policyValue := range policyList {
		if strings.Contains(policyValue[1], in.BucketName) {
			likeBucketList = append(likeBucketList, policyValue[1])
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
			// 保留两位小鼠
			value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", bucketItem.CapacityCur), 64)
			bucketInfo.CapacityCur = decimal.NewFromFloat(value).String()
			bucketInfo.ObjectNum = bucketItem.ObjectNum
			bucketInfo.Permission = bucketItem.Permission
			bucketInfo.CreateTime = bucketItem.CreateTime.Format("2006-01-02 15:04:05")
			bucketInfo.UpdateTime = bucketItem.UpdateTime.Format("2006-01-02 15:04:05")

			resp = append(resp, &bucketInfo)
		}
	}

	return &bucket.AllBucketResp{
		BucketList: resp,
	}, nil
}
