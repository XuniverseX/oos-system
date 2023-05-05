package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"io/ioutil"
	"oos-system/app/rpc/metadata/constent"
	"oos-system/app/rpc/metadata/internal/svc"
	"oos-system/app/rpc/metadata/metaInfo"
	"oos-system/app/rpc/metadata/pb"
	"oos-system/common/xerr"
	"os"
	"path"
	"sync"
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

func (l *GetLogic) Get(in *pb.GetReq) (*pb.GetResp, error) {
	// 拿到桶或者对象对应的读写锁， 不存在就创建
	var rwLock sync.RWMutex
	var lock = &rwLock
	if mutex, ok := l.svcCtx.RWLockMap.LoadOrStore(in.BucketName+in.ObjectName, &rwLock); ok {
		lock = mutex.(*sync.RWMutex)
	}

	//判断对象是否存在
	location := path.Join([]string{l.svcCtx.Config.BaseLocation, in.BucketName, in.ObjectName}...)
	_, err := os.Stat(location)
	if err != nil && os.IsNotExist(err) {
		return nil, errors.Wrapf(xerr.NewErrMsg("对象不存在"), "bucketName: %s objectName: %s",
			in.BucketName, in.ObjectName)
	}

	replicationBase := l.svcCtx.Config.ReplicationLocation

	// 读取objectId， 加读锁， 防止写时去读取导致一致性问题， 双检查
	lock.RLock()
	_, err = os.Stat(location)
	if err != nil && os.IsNotExist(err) {
		lock.RUnlock()
		return nil, errors.Wrapf(xerr.NewErrMsg("对象不存在"), "bucketName: %s objectName: %s",
			in.BucketName, in.ObjectName)
	}
	file, _ := ioutil.ReadFile(path.Join(location, constent.MetaFileName))
	objectInfo := &metaInfo.ObjectInfo{}
	jsonx.Unmarshal(file, objectInfo)
	lock.RUnlock()

	// 若当前数据丢失， 使用备份数据替换， 若备份数据也丢失则数据完全丢失
	_, err = os.Stat(path.Join(replicationBase, in.BucketName, in.ObjectName, objectInfo.ObjectId))
	dir, err := ioutil.ReadDir(path.Join(replicationBase, in.BucketName, in.ObjectName, objectInfo.ObjectId))
	if err != nil && (os.IsNotExist(err) || len(dir) == 0) {
		if len(objectInfo.ReplicationAddr) == 0 {
			return nil, errors.Wrapf(xerr.NewErrMsg("数据丢失"), "bucketName: %s objectName: %s",
				in.BucketName, in.ObjectName)
		}
		lock.Lock()
		defer lock.Unlock()
		file, _ := ioutil.ReadFile(path.Join(location, constent.MetaFileName))
		objectInfo := &metaInfo.ObjectInfo{}
		jsonx.Unmarshal(file, objectInfo)
		_, err = os.Stat(path.Join(replicationBase, in.BucketName, in.ObjectName, objectInfo.ObjectId))
		dir, err := ioutil.ReadDir(path.Join(replicationBase, in.BucketName, in.ObjectName, objectInfo.ObjectId))
		if err != nil && (os.IsNotExist(err) || len(dir) == 0) {
			if len(objectInfo.ReplicationAddr) == 0 {
				return nil, errors.Wrapf(xerr.NewErrMsg("数据丢失"), "bucketName: %s objectName: %s",
					in.BucketName, in.ObjectName)
			}
			objectId := objectInfo.ReplicationAddr[len(objectInfo.ReplicationAddr)-1]
			objectInfo.ReplicationAddr = objectInfo.ReplicationAddr[:len(objectInfo.ReplicationAddr)-1]
			objectInfo.ObjectId = objectId
			toString, _ := jsonx.MarshalToString(objectInfo)
			ioutil.WriteFile(path.Join(location, constent.MetaFileName), []byte(toString), os.ModePerm)
		}
	}

	return &pb.GetResp{
		ObjectId: path.Join(replicationBase, in.BucketName, in.ObjectName, objectInfo.ObjectId),
	}, nil
}
