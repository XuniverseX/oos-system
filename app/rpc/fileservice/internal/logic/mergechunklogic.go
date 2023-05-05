package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"oos-system/common/xerr"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"oos-system/app/rpc/fileservice/internal/svc"
	"oos-system/app/rpc/fileservice/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MergeChunkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMergeChunkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MergeChunkLogic {
	return &MergeChunkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//const fileDir = "/Users/xuni/file"

// MergeChunk rpc verify(VerifyReq) returns (VerifyResp);
func (l *MergeChunkLogic) MergeChunk(in *pb.MergeChunkReq) (*pb.MergeChunkResp, error) {
	var wg sync.WaitGroup

	hash := in.Hash                                                                         // md5校验值
	chunkSize := in.ChunkSize                                                               // 切片大小
	writeFilename := in.Filename                                                            // 合并文件的全路径
	tempDir := path.Join(l.svcCtx.Config.TempPath, in.TempPathPrefix, in.OriFilename, hash) // 分片文件夹目录
	// 获得文件夹下所有文件名
	var s []string
	s, err := GetAllFile(tempDir, s)
	if err != nil {
		return &pb.MergeChunkResp{}, errors.Wrap(xerr.NewErrMsg("临时文件获取出错"), fmt.Sprintln(err))
	}

	// 遍历所有临时文件并写入新文件
	maxThread := l.svcCtx.Config.MaxThread
	var count int64 = 0
	for _, filename := range s {
		for count > maxThread {
			time.Sleep(1 * time.Millisecond)
		}
		// 取得切片下标
		split := strings.Split(filename, "-")
		mark := len(split) - 1
		index, _ := strconv.Atoi(split[mark])
		filename1 := filename
		// goroutine写入文件
		wg.Add(1)
		// count原子自增1
		go func() {
			err2 := func(filename string, pos int64) error {
				// waitGroup
				defer wg.Done()
				// 原子减少1
				atomic.AddInt64(&count, 1)
				defer atomic.AddInt64(&count, -1)
				// 打开file文件
				file, err2 := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
				if err2 != nil {
					return errors.Wrap(xerr.NewErrMsg(fmt.Sprintf(filename, "打开出错")), fmt.Sprintln(err))
				}
				defer file.Close()
				// 创建size大小字节缓冲区
				buf := make([]byte, chunkSize*2)

				// 读取文件
				buflen, err1 := file.Read(buf)
				if err1 != nil {
					return errors.Wrap(xerr.NewErrMsg(fmt.Sprintf(filename, "读取出错")), fmt.Sprintln(err))
				}

				//解压数据
				//debuff, _ := common.GzipDecode(buf)
				f, err1 := os.OpenFile(writeFilename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
				if err1 != nil {
					return errors.Wrap(xerr.NewErrMsg(fmt.Sprintf(writeFilename, "打开出错")), fmt.Sprintln(err))
				}
				defer f.Close()
				_, err1 = f.Seek(pos, io.SeekStart)
				if err1 != nil {
					return errors.Wrap(xerr.NewErrMsg(fmt.Sprintf(writeFilename, "光标移动出错")), fmt.Sprintln(err))
				}
				_, err1 = f.Write(buf[:buflen])
				if err1 != nil {
					return errors.Wrap(xerr.NewErrMsg(fmt.Sprintf(writeFilename, "写入出错")), fmt.Sprintln(err))
				}
				return nil
			}(filename1, chunkSize*int64(index-1))
			if err2 != nil {
				//panic(err2)
			}
		}()
	}
	wg.Wait()

	// todo:md5校验

	return &pb.MergeChunkResp{Filename: writeFilename}, nil
}
