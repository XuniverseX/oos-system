package file

import (
	"context"
	"fmt"
	"oos-system/app/rpc/metadata/metadataservice"
	"path"

	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadUrlLogic {
	return &DownloadUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadUrlLogic) DownloadUrl(req *types.DownloadUrlReq) (resp *types.DownloadUrlResp, err error) {
	// todo: add your logic here and delete this line
	//api服务地址
	url := l.svcCtx.Config.Host
	bucketName := path.Join(req.BucketName, req.Path)
	objectName := req.ObjectName
	metadataRes, err := l.svcCtx.MetadataRpc.Get(l.ctx, &metadataservice.GetReq{
		BucketName: bucketName,
		ObjectName: objectName,
	})
	//
	objId := metadataRes.ObjectId
	objPath := path.Join(objId, objectName)
	//fmt.Println(objId)
	realUrl := path.Join(url, fmt.Sprintf("zhiyi-obj?path=%s", objPath))
	return &types.DownloadUrlResp{
		Success: true,
		Url:     realUrl,
	}, nil
}
