package file

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"
	"oos-system/app/rpc/bucket/bucket"
	"oos-system/app/rpc/fileservice/fileservice"
	"oos-system/common/ctxdata"
	"oos-system/common/xerr"
	"path"
)

type UploadChunkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadChunkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadChunkLogic {
	return &UploadChunkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadChunkLogic) UploadChunk(req *types.UploadChunkReq, r *http.Request) (resp *types.UploadChunkResp, err error) {
	username := ctxdata.GetUserNameFromCtx(l.ctx)

	userPolicy, policyErr := l.svcCtx.BucketRpc.GetPolicy(l.ctx, &bucket.GetPolicyReq{
		Username:   username,
		BucketName: req.BucketName,
	})

	//fmt.Println(username, userPolicy.UserPermission)

	if policyErr != nil {
		return nil, policyErr
	}

	// 0 读写 1 只读 2 只写 3 桶拥有者
	if userPolicy.UserPermission == 1 {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(xerr.No_Bucket_Permission, "没有上传文件权限"), "没有上传文件权限 username : %s, bucket: %s", username, req.BucketName)
	}

	hash := req.Identifier
	index := req.ChunkNumber

	//file, err := os.Open("/Users/xuni/file/three")
	file, handle, _ := r.FormFile("file")
	bytes := make([]byte, handle.Size)

	len1, _ := file.Read(bytes)
	state, err := l.svcCtx.FileserviceRpc.UploadChunk(l.ctx, &fileservice.UploadChunkReq{
		Hash:           hash,
		File:           bytes[:len1],
		Index:          index,
		TotalChunk:     req.TotalChunks,
		TempPathPrefix: path.Join(req.BucketName, req.FilePath),
		Filename:       req.Filename,
	})
	if err != nil {
		return &types.UploadChunkResp{}, err
	}

	return &types.UploadChunkResp{
		Uploaded:   state.Indexes,
		NeedMerge:  state.NeedMerge,
		Success:    state.Success,
		SkipUpload: false,
	}, nil
}
