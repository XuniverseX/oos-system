package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"oos-system/app/rpc/metadata/internal/svc"
	"oos-system/app/rpc/metadata/pb"
	"os"
	"path"
	"sync"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLogic) Delete(in *pb.DeleteReq) (*pb.DeleteResp, error) {
	// 拿到桶或者对象对应的读写锁， 不存在就创建
	var rwLock sync.RWMutex
	var lock = &rwLock
	if mutex, ok := l.svcCtx.RWLockMap.LoadOrStore(in.BucketName+in.ObjectName, &rwLock); ok {
		lock = mutex.(*sync.RWMutex)
	}

	// 判断是否已经删除
	location := path.Join([]string{l.svcCtx.Config.BaseLocation, in.BucketName, in.ObjectName}...)
	replication := path.Join([]string{l.svcCtx.Config.ReplicationLocation, in.BucketName, in.ObjectName}...)
	_, err := os.Stat(location)
	if err != nil && os.IsNotExist(err) {
		return &pb.DeleteResp{
			Deleted: true,
		}, nil
	}

	lock.Lock()
	defer lock.Unlock()
	os.RemoveAll(location)
	os.RemoveAll(replication)
	return &pb.DeleteResp{
		Deleted: true,
	}, nil
}
