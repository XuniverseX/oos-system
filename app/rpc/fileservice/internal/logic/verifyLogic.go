package logic

import (
	"context"
	"os"
	"path"
	"strconv"
	"strings"

	"oos-system/app/rpc/fileservice/internal/svc"
	"oos-system/app/rpc/fileservice/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyLogic {
	return &VerifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//const DirName = "chunkDir_"

func (l *VerifyLogic) Verify(in *pb.VerifyReq) (*pb.VerifyResp, error) {
	location := l.svcCtx.Config.TempPath
	hash := in.Hash

	tempDir := path.Join(location, in.TempPathPrefix, in.Filename, hash)

	//获取tempDir文件夹中所有文件名
	var s []string
	var indexes []int64
	s, err := GetAllFile(tempDir, s)
	if err != nil {
		return &pb.VerifyResp{Hash: hash, Indexes: indexes}, nil
	}
	indexes = make([]int64, len(s))
	// 将文件名填入filenames切片中
	for i, str := range s {
		temp := strings.Split(str, string(os.PathSeparator))
		split := strings.Split(temp[len(temp)-1], "-")
		tmpIndex, _ := strconv.Atoi(split[len(split)-1])
		indexes[i] = int64(tmpIndex)
	}
	return &pb.VerifyResp{Hash: hash, Indexes: indexes}, nil
}
