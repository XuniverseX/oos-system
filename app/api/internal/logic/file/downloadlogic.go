package file

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"oos-system/app/rpc/metadata/metadataservice"
	"oos-system/common/xerr"
	"os"
	"path"

	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadLogic {
	return &DownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadLogic) Download(req *types.DownloadReq, w http.ResponseWriter) error {
	// todo: add your logic here and delete this line

	bucketName := path.Join(req.BucketName, req.Path)
	objectName := req.ObjectName
	metadataRes, err := l.svcCtx.MetadataRpc.Get(l.ctx, &metadataservice.GetReq{
		BucketName: bucketName,
		ObjectName: objectName,
	})
	open, err := os.Open(path.Join(metadataRes.ObjectId, objectName))
	defer open.Close()
	if err != nil {
		fmt.Printf("文件 %s 不存在！\n", path.Join(metadataRes.ObjectId, objectName))
		return errors.Wrapf(xerr.NewErrMsg("无效路径~ "), "err: %+v", err)
	}
	io.Copy(w, open)
	return nil
}
