package file

import (
	"context"
	"net/http"
	"oos-system/app/rpc/fileservice/fileservice"
	"oos-system/app/rpc/metadata/metadataservice"

	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadNoChunkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadNoChunkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadNoChunkLogic {
	return &UploadNoChunkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//不分片上传
func (l *UploadNoChunkLogic) UploadNoChunk(req *types.UploadNoChunkReq, r *http.Request) (resp *types.UploadNoChunkResp, err error) {
	// todo: add your logic here and delete this line
	file, handle, _ := r.FormFile("obj")
	bytes := make([]byte, handle.Size)
	//不分片上传直接先获取三个文件地址，循环写入
	state, _ := l.svcCtx.MetadataRpc.Put(l.ctx, &metadataservice.PutReq{
		BucketName: req.BucketName,
		ObjectName: req.ObjectName,
		IsDir:      req.IsDir,
		CreateUser: req.CreateUser,
		CheckCode:  req.CheckCode,
		Size:       req.TotalSize,
	})

	paths := state.ReplicationId
	len1, _ := file.Read(bytes)
	state2, _ := l.svcCtx.FileserviceRpc.Put(l.ctx, &fileservice.PutRequest{
		Files:         bytes[:len1],
		Hash:          req.Hash,
		ReplicationId: paths,
	})

	return &types.UploadNoChunkResp{
		Success: state2.Success,
	}, nil
}
