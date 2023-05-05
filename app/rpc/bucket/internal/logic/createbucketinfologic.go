package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/bucket/internal/svc"
	"oos-system/app/rpc/model/bucketmodel"
	"oos-system/common/xerr"
)

type CreateBucketInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBucketInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBucketInfoLogic {
	return &CreateBucketInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateBucketInfoLogic) CreateBucketInfo(in *bucket.CreateBucketInfoReq) (*bucket.SucResp, error) {
	// 写入bucket表 和 权限表 使用本地事务
	// casbin事务没法回滚 还得自己导入casbin的model 所以这块为了方便自己手写事务 第二个表失败直接把第一个表删了就行

	// 插入bucket表
	_, err := l.svcCtx.BucketModel.Insert(l.ctx, &bucketmodel.Bucket{
		Username:    in.Username,
		BucketName:  in.BucketName,
		CapacityCur: float64(in.CapacityCur),
		ObjectNum:   in.ObjectNum,
		Permission:  in.Permission,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "insert bucketInfo db Insert err:%v,username:%+v", err, in.Username)
	}
	// 添加权限表 自己的桶拥有读写权限
	// 0 读写权限 1 读权限 2 写权限 3 表示桶的拥有者
	_, casbinErr := l.svcCtx.Casbin.AddPolicy(in.Username, in.BucketName, "3")
	if casbinErr != nil {
		// 简易实现 自己手写一个假装事务
		// 插入权限表失败了就把bucket删了
		delBucketErr := l.svcCtx.BucketModel.DeleteByBucketName(l.ctx, in.BucketName)
		if delBucketErr != nil {
			return nil, delBucketErr
		}
		return nil, errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "casbin 新建桶 桶拥有者权限添加失败")
	}

	return &bucket.SucResp{
		Success: true,
	}, nil
}
