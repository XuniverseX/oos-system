package logic

import (
	"bufio"
	"context"
	"fmt"
	"oos-system/app/rpc/fileservice/internal/svc"
	"oos-system/app/rpc/fileservice/pb"
	"os"

	"github.com/zeromicro/go-zero/core/logx"
)

type PreviewLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPreviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PreviewLogic {
	return &PreviewLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PreviewLogic) Preview(in *pb.PreviewReq) (*pb.PreviewResp, error) {
	// todo: add your logic here and delete this line

	//wtFilename := fmt.Sprintf("%s%s%s", l.svcCtx.Config.WriteToPath, string(os.PathSeparator), in.Hash)
	wtFilename := ""

	file, _ := os.Open(wtFilename)
	stat, _ := file.Stat()

	source := make([]byte, stat.Size())
	fmt.Println("打开：", wtFilename)
	defer file.Close()
	bufio.NewReader(file).Read(source)
	return &pb.PreviewResp{
		Data: source,
	}, nil
	return &pb.PreviewResp{}, nil
}
