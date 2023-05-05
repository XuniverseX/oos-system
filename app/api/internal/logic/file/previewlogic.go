package file

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"
	"oos-system/common/xerr"
	"os"
)

type PreviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPreviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PreviewLogic {
	return &PreviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//预览||下载
func (l *PreviewLogic) Preview(req *types.PreviewReq, w http.ResponseWriter, r *http.Request) error {
	// todo: add your logic here and delete this line
	objPath := req.Path
	open, err := os.Open(objPath)
	defer open.Close()
	if err != nil {
		fmt.Printf("文件 %s 不存在！\n", objPath)
		return errors.Wrapf(xerr.NewErrMsg("无效路径~ "), "err: %+v", err)
	}
	io.Copy(w, open)
	return nil
}
