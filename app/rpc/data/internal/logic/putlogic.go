package logic

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"oos-system/app/rpc/data/internal/svc"
	"oos-system/app/rpc/data/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutLogic {
	return &PutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/**
index 分片索引
totalBlock 总分片数量
*/
func (l *PutLogic) Put(in *pb.PutRequest) (*pb.PutResponse, error) {
	// todo: add your logic here and delete this line
	var sourcedata = in.GetFiles()
	fmt.Println(len(sourcedata), "============= 索引 :")

	//root/hash/index
	objRealPath := fmt.Sprintf("%s%s%s", l.svcCtx.Config.RootPath, string(os.PathSeparator), in.Hash)
	if _, err := os.Stat(objRealPath); os.IsNotExist(err) {
		os.MkdirAll(objRealPath, os.ModePerm)
	}
	sli := fmt.Sprintf("%s%s%s-%s", objRealPath, string(os.PathSeparator), in.Hash, strconv.Itoa(int(in.Index)))
	obj, _ := os.Create(sli)
	defer obj.Close()
	obj.Write(sourcedata)
	return &pb.PutResponse{Path: objRealPath}, nil
}
