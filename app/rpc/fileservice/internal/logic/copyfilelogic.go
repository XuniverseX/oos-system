package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"oos-system/app/rpc/fileservice/internal/svc"
	"oos-system/app/rpc/fileservice/pb"
	"oos-system/common/xerr"
	"os"
	"path"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type CopyFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCopyFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CopyFileLogic {
	return &CopyFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CopyFileLogic) CopyFile(in *pb.CopyFileReq) (*pb.CopyFileResp, error) {
	url := in.OriginFilename    // 源文件url
	paths := in.ReplicationPath // 副本目录

	//split := strings.Split(url, string(os.PathSeparator))
	split := strings.Split(url, "/")
	filename := split[len(split)-1]

	for _, pathAddr := range paths {
		join := path.Join(pathAddr, filename)
		read, err := os.OpenFile(url, os.O_RDONLY, os.ModePerm)
		if err != nil {
			return &pb.CopyFileResp{Success: false}, errors.Wrap(xerr.NewErrMsg("文件打开失败"), fmt.Sprintln(url, "文件打开失败"))
		}
		write, err := os.OpenFile(join, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			return &pb.CopyFileResp{Success: false}, errors.Wrap(xerr.NewErrMsg("文件打开失败"), fmt.Sprintln(join, "文件打开失败"))
		}
		_, err = io.Copy(write, read)
		err = read.Close()
		err = write.Close()
		if err != nil {
			return &pb.CopyFileResp{Success: false}, errors.Wrap(xerr.NewErrMsg("文件复制出错"), fmt.Sprintln("文件复制出错"))
		}
	}

	return &pb.CopyFileResp{Success: true}, nil
}
