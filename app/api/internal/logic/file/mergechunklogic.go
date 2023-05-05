package file

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"
	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/fileservice/fileservice"
	"oos-system/app/rpc/metadata/pb"
	"oos-system/common/ctxdata"
	"os"
	"path"
	"strconv"
)

type MergeChunkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMergeChunkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MergeChunkLogic {
	return &MergeChunkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MergeChunkLogic) MergeChunk(req *types.MergeChunkReq) (resp *types.MergeChuckResp, err error) {

	hashCode := req.HashCode
	chunkSize, _ := strconv.ParseInt(req.ChunkSize, 10, 64)

	//fmt.Println("hashTest: ", req.FileName, hashCode)

	//分片上传完成后， 调用 metadata.put获取三个文件夹地址
	putResp, err := l.svcCtx.MetadataRpc.Put(l.ctx, &pb.PutReq{
		BucketName: req.BucketName,
		ObjectName: path.Join(req.FilePath, req.FileName),
		IsDir:      false,
		CreateUser: ctxdata.GetUserNameFromCtx(l.ctx),
		CheckCode:  hashCode,
		Size:       req.TotalSize,
	})

	//merge的时候写入新文件 把哈希写入数据库
	_, errRpcInsert := l.svcCtx.BucketRpc.AddHashCode(l.ctx, &bucket.HashCodeReq{
		Hashcode: path.Join(req.BucketName, req.FilePath, req.FileName, hashCode),
	})
	if errRpcInsert != nil {
		return &types.MergeChuckResp{Success: false}, nil
	}

	// 将数据库中桶信息更新 桶大小更新 对象数++
	_, updateBucketErr := l.svcCtx.BucketRpc.UpdateBucketSizeAndNumInfo(l.ctx, &bucket.UpdateBucketSizeAndNumReq{
		BucketName: req.BucketName,
		Size:       req.TotalSize,
		ObjectNum:  1,
	})
	if updateBucketErr != nil {
		return nil, updateBucketErr
	}
	//fmt.Println("updateDbBucket:", updateDb.Success)

	// 主目录写入
	_, err = l.svcCtx.FileserviceRpc.MergeChunk(l.ctx, &fileservice.MergeChunkReq{
		Filename:       path.Join(putResp.ObjectId, req.FileName),
		Hash:           hashCode,
		ChunkSize:      chunkSize,
		TempPathPrefix: path.Join(req.BucketName, req.FilePath),
		OriFilename:    req.FileName,
	})
	if err != nil {
		return &types.MergeChuckResp{Success: false}, err
	}

	//for _, replication := range putResp.ReplicationId {
	//	_, err = l.svcCtx.FileserviceRpc.MergeChunk(l.ctx, &fileservice.MergeChunkReq{
	//		Filename:       path.Join(replication, req.FileName),
	//		Hash:           hashCode,
	//		ChunkSize:      chunkSize,
	//		TempPathPrefix: path.Join(req.BucketName, req.FilePath),
	//		OriFilename:    req.FileName,
	//	})
	//}
	//压缩
	//l.svcCtx.FileserviceRpc.Compression(l.ctx, &fileservice.CompressionReq{
	//	Filename: path.Join(putResp.ReplicationId[0], req.FileName),
	//	Hash:     hashCode,
	//})

	//删除切片文件夹
	err = os.RemoveAll(path.Join(l.svcCtx.Config.TempPath, req.BucketName, req.FilePath, req.FileName))
	if err != nil {
		return &types.MergeChuckResp{Success: false}, err
	}

	//复制到副本目录
	//_, err = l.svcCtx.FileserviceRpc.CopyFile(l.ctx, &fileservice.CopyFileReq{
	//	OriginFilename:  mergeChunkResp.Filename,
	//	ReplicationPath: putResp.ReplicationId,
	//})
	//if err != nil {
	//	return &types.MergeChuckResp{Success: false}, err
	//}

	return &types.MergeChuckResp{Success: true}, nil
}
