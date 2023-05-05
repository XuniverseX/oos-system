package logic

import (
	"bufio"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"oos-system/app/rpc/fileservice/internal/svc"
	"oos-system/app/rpc/fileservice/pb"
	"oos-system/common/xerr"
	"os"
	"path"
	"strconv"
	"strings"
)

type UploadChunkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadChunkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadChunkLogic {
	return &UploadChunkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadChunkLogic) UploadChunk(in *pb.UploadChunkReq) (*pb.UploadChunkResp, error) {
	b := in.GetFile()
	hash := in.Hash
	index := in.GetIndex()

	//文件名与路径
	tempDir := path.Join(l.svcCtx.Config.TempPath, in.TempPathPrefix, in.Filename, hash)
	filename := fmt.Sprintf("%s/%s-%s", tempDir, hash, strconv.Itoa(int(index)))

	_, err := os.Stat(tempDir)
	if os.IsNotExist(err) {
		//创建临时目录，tempDir+hash为目录名
		err1 := os.MkdirAll(tempDir, os.ModePerm)
		if err1 != nil {
			return &pb.UploadChunkResp{Success: false}, errors.Wrap(xerr.NewErrMsg("文件目录创建失败"), fmt.Sprintf(tempDir, "目录创建失败 err: %v", err))
		}
	}
	//创建文件
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	//f, err := os.Create(filename)
	if err != nil {
		return &pb.UploadChunkResp{Success: false}, errors.Wrap(xerr.NewErrMsg("文件打开失败"), fmt.Sprintf(filename, "打开失败 err: %v", err))
	}
	defer f.Close()
	writer := bufio.NewWriterSize(f, 2*1024*1024)
	_, err = writer.Write(b)
	if err != nil {
		return &pb.UploadChunkResp{Success: false}, errors.Wrap(xerr.NewErrMsg("文件写入缓冲出错"), fmt.Sprintf(filename, "写入缓冲出错 err: %v", err))
	}
	err = writer.Flush()
	if err != nil {
		return &pb.UploadChunkResp{Success: false}, errors.Wrap(xerr.NewErrMsg("文件刷新缓冲出错"), fmt.Sprintf("刷新缓冲出错 err: %v", err))
	}

	//获取tempDir文件夹中所有文件名
	var s []string
	s, err = GetAllFile(tempDir, s)
	if err != nil {
		return nil, err
	}
	temp := make([]int64, len(s))
	for i, str := range s {
		split := strings.Split(str, "-")
		temp[i], err = strconv.ParseInt(split[len(split)-1], 10, 64)
		if err != nil {
			return &pb.UploadChunkResp{Success: false}, errors.Wrap(xerr.NewErrMsg(fmt.Sprintf(filename, "下标转换int出错")), fmt.Sprintln(err))
		}
	}
	indexes := make([]int64, len(s))
	for i, val := range temp {
		indexes[i] = val
	}

	// 判断是否需要合并
	total := int(in.TotalChunk)
	if err != nil {
		return &pb.UploadChunkResp{Success: false}, errors.Wrap(xerr.NewErrMsg(fmt.Sprintf(filename, "total转换int出错")), fmt.Sprintln(err))
	}
	count := len(s)
	flag := false
	if count == total {
		flag = true
	}

	return &pb.UploadChunkResp{
		Indexes:   indexes,
		Success:   true,
		NeedMerge: flag,
	}, nil
}

// GetAllFile 获取文件夹下所有文件
func GetAllFile(pathname string, s []string) ([]string, error) {
	rd, err := os.ReadDir(pathname)
	if err != nil {
		return s, err
	}

	for _, fi := range rd {
		if !fi.IsDir() {
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}

// GetFileCount 获取tempDir文件夹中文件数量
func GetFileCount(pathname string) (count int, err error) {
	rd, err := os.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return -1, err
	}

	for _, fi := range rd {
		if !fi.IsDir() {
			count++
		}
	}
	return count, nil
}
