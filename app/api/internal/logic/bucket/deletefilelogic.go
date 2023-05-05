package bucket

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"
	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/metadata/pb"
	"oos-system/common/ctxdata"
	"oos-system/common/xerr"
	"path"
)

type DeleteFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFileLogic {
	return &DeleteFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFileLogic) DeleteFile(req *types.DeleteFileReq) (resp *types.DeleteFileResp, err error) {
	username := ctxdata.GetUserNameFromCtx(l.ctx)

	// todo: woc 这块的事务一致性能手写到死

	userPolicy, policyErr := l.svcCtx.BucketRpc.GetPolicy(l.ctx, &bucket.GetPolicyReq{
		Username:   username,
		BucketName: req.BucketName,
	})

	if policyErr != nil {
		return nil, policyErr
	}

	//fmt.Printf("username: %s, policy: %v", username, userPolicy.UserPermission)

	// 0 读写权限 1 只读 2 只写 3 桶拥有者
	// 只读权限没有删除桶中文件权限
	if userPolicy.UserPermission == 1 {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(xerr.No_Bucket_Permission, "没有权限删除桶中文件及文件夹"), "没有权限删除桶中文件及文件夹 username : %s, bucket: %s , permission: %d", username, req.BucketName, userPolicy.UserPermission)
	}

	// 只有桶拥有者才可以删除桶
	// 删除根目录就是删除桶 需要把几个表里的桶全删了
	if req.Path == "/" && userPolicy.UserPermission == 3 {
		policy, err := l.svcCtx.BucketRpc.DelBucketAllPolicy(l.ctx, &bucket.DelBucketAllPolicyReq{
			BucketName: req.BucketName,
		})
		if err != nil {
			return nil, err
		}

		// 然后删除桶数据库  todo 我也不想管事务了md 这要是事务都得手写啊 casbin 又没办法事务
		delBucket, err := l.svcCtx.BucketRpc.DelBucket(l.ctx, &bucket.DelBucketReq{
			BucketName: req.BucketName,
		})

		if err != nil {
			return nil, err
		}

		// 删除哈希表哈希 我表示不想处理这个err了
		_, _ = l.svcCtx.BucketRpc.DelBucketAllHash(l.ctx, &bucket.DelBucketAllHashReq{BucketName: req.BucketName})

		if !(policy.Success && delBucket.Success) {
			return nil, errors.Wrapf(xerr.NewErrMsg("桶删除失败"), "删除桶失败 嘎蛋，会出大事 err : %v", err)
		}
	} else if req.Path == "/" && userPolicy.UserPermission != 3 {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(xerr.No_Bucket_Permission, "没有桶权限删除文件及文件夹权限"), "没有桶权限删除文件及文件夹权限 username : %s, bucket: %s , permission: %d", username, req.BucketName, userPolicy.UserPermission)
	}

	// 是文件夹 不是 桶目录
	if req.IsDir && req.Path != "/" {
		// 删除哈希表里对应文件夹下的hash
		// 直接模糊查询全干完了
		//println(req.BucketName + req.Path)
		_, _ = l.svcCtx.BucketRpc.DelBucketAllHash(l.ctx, &bucket.DelBucketAllHashReq{BucketName: path.Join(req.BucketName, req.Path)})
	}

	// 删除本地文件
	deleteResp, err := l.svcCtx.MetadataRpc.Delete(l.ctx, &pb.DeleteReq{
		ObjectName: req.Path,
		BucketName: req.BucketName,
		IsDir:      req.IsDir,
	})
	if !deleteResp.Deleted {
		return nil, err
	}

	// 不是文件夹再去删除文件需要删除的单个hash
	if !req.IsDir {
		// 删除桶中文件 需要先删除哈希表中文件哈希
		delHash, delHashErr := l.svcCtx.BucketRpc.DelHashCode(l.ctx, &bucket.HashCodeReq{
			Hashcode: path.Join(req.BucketName, req.Path, req.HashCode),
		})
		if delHashErr != nil {
			return nil, delHashErr
		}

		if delHash.Success {
			// 然后再去对桶信息更新
			_, UpdateErr := l.svcCtx.BucketRpc.UpdateBucketSizeAndNumInfo(l.ctx, &bucket.UpdateBucketSizeAndNumReq{
				BucketName: req.BucketName,
				Size:       -req.Size,
				ObjectNum:  -1,
			})
			if UpdateErr != nil {
				return nil, UpdateErr
			}
		}
	}

	return &types.DeleteFileResp{}, nil
}
