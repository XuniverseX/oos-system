package logic

import (
	"context"
	"github.com/google/uuid"
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
	"time"
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

func (l *PutLogic) Put(in *pb.PutReq) (*pb.PutResp, error) {
	// 拿到桶或者对象对应的读写锁， 不存在就创建
	var rwLock sync.RWMutex
	var lock = &rwLock
	preHash := ""
	if mutex, ok := l.svcCtx.RWLockMap.LoadOrStore(in.BucketName+in.ObjectName, &rwLock); ok {
		lock = mutex.(*sync.RWMutex)
	}

	// 对对象是否存在进行判断， 这里需要对对象的校验码进行比较如果一致，则无需上传
	location := path.Join([]string{l.svcCtx.Config.BaseLocation, in.BucketName, in.ObjectName}...)
	replicationBase := l.svcCtx.Config.ReplicationLocation
	_, err := os.Stat(location)
	if err == nil {
		_, err := os.Stat(path.Join(location, constent.MetaFileName))
		if err == nil {
			lock.RLock()
			readFile, _ := ioutil.ReadFile(path.Join(location, constent.MetaFileName))
			lock.RUnlock()
			objectInfo := &metaInfo.ObjectInfo{}
			jsonx.UnmarshalFromString(string(readFile), objectInfo)
			if objectInfo.ChekCode == in.CheckCode {
				objId := objectInfo.ObjectId
				return &pb.PutResp{
					ObjectId:      objId,
					AlreadyUpload: true,
					PreHash:       "",
				}, nil
			}
			preHash = objectInfo.ChekCode
			os.RemoveAll(path.Join(location, constent.MetaFileName))
			os.RemoveAll(path.Join(l.svcCtx.Config.ReplicationLocation, in.BucketName, in.ObjectName))
		}
	} else if !os.IsNotExist(err) {
		return nil, errors.Wrapf(xerr.NewErrMsg("系统错误"), "err: %+v", err)
	}

	// 创建对象信息结构体， 若对象为目录类型的， 则objectId就是对象名， 否则使用uuid生成唯一对象名
	objectInfo := &metaInfo.ObjectInfo{
		ObjectName: in.ObjectName,
		ObjectId:   uuid.New().String(),
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		CreateUser: in.CreateUser,
		IsDir:      in.IsDir,
		ChekCode:   in.CheckCode,
		Size:       in.Size,
	}
	objId := objectInfo.ObjectId
	if !objectInfo.IsDir {
		objectInfo.ReplicationAddr = []string{uuid.NewString(), uuid.NewString()}
	}
	toString, _ := jsonx.MarshalToString(objectInfo)

	// 加写锁， 进行对象元数据写入， 双检查
	lock.Lock()
	defer lock.Unlock()
	_, err = os.Stat(location)
	if err == nil {
		_, err := os.Stat(path.Join(location, constent.MetaFileName))
		if err == nil {
			readFile, _ := ioutil.ReadFile(path.Join(location, constent.MetaFileName))
			jsonx.UnmarshalFromString(string(readFile), objectInfo)
			objId = objectInfo.ObjectId
			if objectInfo.ChekCode == in.CheckCode {
				return &pb.PutResp{
					ObjectId:      objId,
					AlreadyUpload: true,
				}, nil
			}
			preHash = objectInfo.ChekCode
			os.RemoveAll(path.Join(location, constent.MetaFileName))
			os.RemoveAll(path.Join(l.svcCtx.Config.ReplicationLocation, in.BucketName, in.ObjectName))
		}
	} else if !os.IsNotExist(err) {
		return nil, errors.Wrapf(xerr.NewErrMsg("系统错误"), "err: %+v", err)
	}

	os.Mkdir(location, os.ModePerm)
	// 若不为文件夹则创建以ObjectId为名的文件夹，文件夹下存放真实数据
	if !in.IsDir {
		os.MkdirAll(path.Join(replicationBase, in.BucketName, in.ObjectName, objectInfo.ObjectId), os.ModePerm)
		for _, v := range objectInfo.ReplicationAddr {
			os.MkdirAll(path.Join(replicationBase, in.BucketName, in.ObjectName, v), os.ModePerm)
		}
	}
	// 元数据写入
	ioutil.WriteFile(path.Join(location, constent.MetaFileName), []byte(toString), os.ModePerm)

	addr := make([]string, 0)
	for _, v := range objectInfo.ReplicationAddr {
		addr = append(addr, path.Join(replicationBase, in.BucketName, in.ObjectName, v))
	}
	return &pb.PutResp{
		ObjectId:      path.Join(replicationBase, in.BucketName, in.ObjectName, objectInfo.ObjectId),
		AlreadyUpload: false,
		ReplicationId: addr,
		PreHash:       preHash,
	}, nil
}
