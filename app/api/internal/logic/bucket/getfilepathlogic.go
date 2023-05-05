package bucket

import (
	"context"
	"github.com/pkg/errors"
	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"
	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/metadata/pb"
	"oos-system/common/ctxdata"
	"oos-system/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFilePathLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFilePathLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFilePathLogic {
	return &GetFilePathLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFilePathLogic) GetFilePath(req *types.FilePathReq) (resp *types.FilePathResp, err error) {
	username := ctxdata.GetUserNameFromCtx(l.ctx)

	userPolicy, policyErr := l.svcCtx.BucketRpc.GetPolicy(l.ctx, &bucket.GetPolicyReq{
		Username:   username,
		BucketName: req.BucketName,
	})

	if policyErr != nil {
		return nil, policyErr
	}

	if userPolicy.UserPermission == -1 {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(xerr.No_Bucket_Permission, "没有获取桶列表权限"), "没有获取桶文件目录权限 username : %s, bucket: %s", username, req.BucketName)
	}

	filePath, err := l.svcCtx.MetadataRpc.BucketInfo(l.ctx, &pb.BucketInfoReq{
		BucketName: req.BucketName,
		ObjectName: req.Path,
	})
	//fmt.Printf("filepath: %v \n", filePath)
	if err != nil {
		return nil, err
	}

	var filePathList []types.ObjectInfo

	if len(filePath.ObjectInfo) > 0 {
		for _, item := range filePath.ObjectInfo {
			var typesFilePath types.ObjectInfo

			typesFilePath.Name = item.Name
			typesFilePath.ObjectName = item.ObjectName
			typesFilePath.CreateTime = item.CreateTime
			typesFilePath.Size = item.Size
			typesFilePath.IsDir = item.IsDir
			typesFilePath.HashCode = item.Hash

			filePathList = append(filePathList, typesFilePath)
		}
	}

	return &types.FilePathResp{
		BucketPathList: filePathList,
	}, nil
}
