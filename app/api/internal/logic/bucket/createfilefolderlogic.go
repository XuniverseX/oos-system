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

type CreateFileFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateFileFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFileFolderLogic {
	return &CreateFileFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateFileFolderLogic) CreateFileFolder(req *types.CreateFileFolderReq) (resp *types.CreateFileFolderResp, err error) {
	username := ctxdata.GetUserNameFromCtx(l.ctx)

	policyRes, getPolicyErr := l.svcCtx.BucketRpc.GetPolicy(l.ctx, &bucket.GetPolicyReq{
		Username:   username,
		BucketName: req.BucketName,
	})

	if getPolicyErr != nil {
		return nil, getPolicyErr
	}
	// 0读写 1只读 2只写 3 拥有者
	// 桶中创建文件夹
	if policyRes.UserPermission == 1 {
		return nil, errors.Wrapf(xerr.NewErrMsg("创建文件夹没有权限"), "创建文件夹没有桶权限")
	}

	_, err = l.svcCtx.MetadataRpc.Put(l.ctx, &pb.PutReq{
		BucketName: req.BucketName,
		ObjectName: req.FileFloderPath,
		IsDir:      true,
		CreateUser: username,
		CheckCode:  "",
		Size:       0,
	})

	if err != nil {
		return nil, err
	}

	return &types.CreateFileFolderResp{}, nil
}
