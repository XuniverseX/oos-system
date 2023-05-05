package logic

import (
	"context"
	"fmt"
	"oos-system/app/rpc/fileservice/internal/svc"
	"oos-system/app/rpc/fileservice/pb"
	"os"

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

//不做分片的文件就调用这个
func (l *PutLogic) Put(in *pb.PutRequest) (*pb.PutResponse, error) {
	// todo: add your logic here and delete this line
	var sourcedata = in.GetFiles()

	//root/hash/index
	//objRealPath := fmt.Sprintf("%s%s%s", l.svcCtx.Config.TempPath, string(os.PathSeparator), in.Hash)
	//if _, err := os.Stat(objRealPath); os.IsNotExist(err) {
	//	os.MkdirAll(objRealPath, os.ModePerm)
	//}
	//写入三个文件夹
	paths := in.ReplicationId
	for _, path := range paths {
		objRealPath := fmt.Sprintf("%s%s%s", path, string(os.PathSeparator), in.Hash)
		obj, _ := os.Create(objRealPath)
		defer obj.Close()
		obj.Write(sourcedata)
	}

	return &pb.PutResponse{Success: true}, nil
}
