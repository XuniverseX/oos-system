package bucket

import (
	"context"
	"github.com/pkg/errors"
	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/metadata/pb"
	"oos-system/common/ctxdata"
	"oos-system/common/xerr"

	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBucketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateBucketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBucketLogic {
	return &CreateBucketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateBucketLogic) CreateBucket(req *types.CreateBucketReq) (resp *types.CreateBucketResq, err error) {
	username := ctxdata.GetUserNameFromCtx(l.ctx)

	// 查询用户桶数量是否超过10个
	countBucket, err := l.svcCtx.BucketRpc.CountBucket(l.ctx, &bucket.CountBucketReq{
		Username: username,
	})

	if err != nil {
		return nil, err
	}

	if countBucket.BucketNum >= 1050 {
		return nil, errors.Wrapf(xerr.NewErrMsg("创建桶数量不能超过十个"), "创建桶数量不能超过十个 err: %v, username: %v", err, username)
	}

	// metadata桶创建rpc
	_, err = l.svcCtx.MetadataRpc.Create(l.ctx, &pb.CreateReq{
		Username:   username,
		BucketName: req.BucketName,
		Version:    req.Version,
	})
	if err != nil {
		return nil, err
	}

	// 将桶信息写入数据库 初始化创建的时候 CapacityCur  ObjectNum 都先为 0
	// 同时将对自己的桶权限写入权限表，事务在rpc中处理
	// rpc中处理事务casbin api 无法回滚 所以只能自己手动判断事务
	_, err = l.svcCtx.BucketRpc.CreateBucketInfo(l.ctx, &bucket.CreateBucketInfoReq{
		Username:    username,
		BucketName:  req.BucketName,
		CapacityCur: 0,
		ObjectNum:   0,
		Permission:  req.Permission,
	})
	if err != nil {
		// 事务处理 下面两个表任意一个出了问题都会返回err rpc返回的err在api中捕获然后把桶文件夹删除
		_, delBucketErr := l.svcCtx.MetadataRpc.Delete(l.ctx, &pb.DeleteReq{
			ObjectName: "",
			BucketName: req.BucketName,
			IsDir:      false,
		})
		if delBucketErr != nil {
			return nil, delBucketErr
		}
		return nil, err
	}

	return &types.CreateBucketResq{}, nil
}
