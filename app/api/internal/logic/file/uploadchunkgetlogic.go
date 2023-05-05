package file

import (
	"context"
	"github.com/pkg/errors"
	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/fileservice/fileservice"
	"oos-system/common/ctxdata"
	"oos-system/common/xerr"
	"path"

	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadChunkGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadChunkGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadChunkGetLogic {
	return &UploadChunkGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadChunkGetLogic) UploadChunkGet(req *types.UploadChunkGetReq) (resp *types.UploadChunkResp, err error) {
	// 权限校验
	username := ctxdata.GetUserNameFromCtx(l.ctx)

	userPolicy, policyErr := l.svcCtx.BucketRpc.GetPolicy(l.ctx, &bucket.GetPolicyReq{
		Username:   username,
		BucketName: req.BucketName,
	})

	if policyErr != nil {
		return nil, policyErr
	}

	// 0 读写 1 只读 2 只写 3 桶拥有者
	if userPolicy.UserPermission == 1 {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(xerr.No_Bucket_Permission, "没有上传文件权限"), "没有上传文件权限 username : %s, bucket: %s", username, req.BucketName)
	}
	//先查库 看看是不是秒传
	//fmt.Println(path.Join(req.BucketName, req.FilePath, req.Filename, req.Identifier))
	hashCode, errRpc := l.svcCtx.BucketRpc.HasHashCode(l.ctx, &bucket.HashCodeReq{
		Hashcode: path.Join(req.BucketName, req.FilePath, req.Filename, req.Identifier),
	})
	if errRpc != nil {
		return &types.UploadChunkResp{}, errRpc
	}
	if hashCode.Success {
		var emptyArr []int64
		return &types.UploadChunkResp{
			Uploaded:   emptyArr,
			NeedMerge:  false,
			Success:    hashCode.Success,
			SkipUpload: true,
		}, nil
	}

	verifyResp, err := l.svcCtx.FileserviceRpc.Verify(l.ctx, &fileservice.VerifyReq{
		Hash:           req.Identifier,
		TempPathPrefix: path.Join(req.BucketName, req.FilePath),
		Filename:       req.Filename,
	})
	if err != nil {
		return &types.UploadChunkResp{}, nil
	}

	return &types.UploadChunkResp{
		Uploaded:   verifyResp.Indexes,
		NeedMerge:  false,
		Success:    false,
		SkipUpload: false,
	}, nil
}
