package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/usercenter/usercenter"
	"oos-system/common/ctxdata"
	"oos-system/common/xerr"

	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBySelfLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBySelfLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBySelfLogic {
	return &DeleteBySelfLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteBySelfLogic) DeleteBySelf(req *types.DelReq) (resp *types.DelResp, err error) {
	username := ctxdata.GetUserNameFromCtx(l.ctx)
	// 查一下库 看一下还有没有桶
	countBucket, countBucketErr := l.svcCtx.BucketRpc.CountBucket(l.ctx, &bucket.CountBucketReq{
		Username: username,
	})
	if countBucketErr != nil {
		return nil, countBucketErr
	}
	//fmt.Println(countBucket.BucketNum)

	if countBucket.BucketNum != 0 {
		return nil, errors.Wrapf(xerr.NewErrMsg("还有存在的桶，不允许注销账户"), "用户还存在桶， 不能注销： username: %s, bucketNum: %d", username, countBucket.BucketNum)
	}

	// 删除该用户所有权限
	_, delUserAllPolicy := l.svcCtx.BucketRpc.DelUserAllPolicy(l.ctx, &bucket.DelUserAllPolicy{UserName: username})

	if delUserAllPolicy != nil {
		return nil, delUserAllPolicy
	}

	// 删除用户 直接删除库 不做逻辑删除了
	deleteBySelf, err := l.svcCtx.UsercenterRpc.DeleteBySelf(l.ctx, &usercenter.DelReq{
		Username: username,
	})
	if err != nil {
		return nil, err
	}

	var res types.DelResp
	_ = copier.Copy(&res, deleteBySelf)

	return &res, nil
	//return &types.DelResp{}, nil
}
