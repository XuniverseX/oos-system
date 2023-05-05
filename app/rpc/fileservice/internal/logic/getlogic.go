package logic

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"

	"oos-system/app/rpc/fileservice/internal/svc"
	"oos-system/app/rpc/fileservice/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogic) Get(in *pb.GetRequest) (*pb.GetResponse, error) {
	// todo: add your logic here and delete this line
	objRealPath := fmt.Sprintf("%s%s%s", l.svcCtx.Config.TempPath, string(os.PathSeparator), in.Hash)
	if _, err := os.Stat(objRealPath); os.IsNotExist(err) {
		return nil, err
	}
	sli := fmt.Sprintf("%s%s%s-%s", objRealPath, string(os.PathSeparator), in.Hash, strconv.Itoa(int(in.Index)))

	file, _ := os.Open(sli)
	stat, _ := file.Stat()

	source := make([]byte, stat.Size())
	fmt.Println("打开：", sli)
	defer file.Close()
	bufio.NewReader(file).Read(source)
	fmt.Println("读取长度：", len(source), ", 分片索引：", in.Index)
	return &pb.GetResponse{
		Data: source,
	}, nil
}
