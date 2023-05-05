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
	"time"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *pb.CreateReq) (*pb.CreateResp, error) {
	// 拿到桶或者对象对应的读写锁， 不存在就创建
	var rwLock sync.RWMutex
	var lock = &rwLock
	if mutex, ok := l.svcCtx.RWLockMap.LoadOrStore(in.BucketName, &rwLock); ok {
		lock = mutex.(*sync.RWMutex)
	}

	// 对桶是否存在进行判断
	location := path.Join([]string{l.svcCtx.Config.BaseLocation, in.BucketName}...)
	_, err := os.Stat(location)
	if err == nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("桶已经存在"), "bucketName: %s", in.BucketName)
	} else if !os.IsNotExist(err) {
		return nil, errors.Wrapf(xerr.NewErrMsg("系统错误"), "err: %+v", err)
	}

	// 创建桶信息结构体， 并序列化便于写入文件
	bucketInfo := &metaInfo.BucketInfo{
		BucketName: in.BucketName,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		CreateUser: in.Username,
		Version:    in.Version,
	}
	toString, _ := jsonx.MarshalToString(bucketInfo)

	// 加写锁， 防止并发问题， 开始写入桶的元数据并对目录进行创建， 双检查
	lock.Lock()
	defer lock.Unlock()
	_, err = os.Stat(location)
	if err == nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("桶已经存在"), "bucketName: %s", in.BucketName)
	} else if !os.IsNotExist(err) {
		return nil, errors.Wrapf(xerr.NewErrMsg("系统错误"), "err: %+v", err)
	}
	os.Mkdir(location, os.ModePerm)
	ioutil.WriteFile(path.Join(location, constent.MetaFileName), []byte(toString), os.ModePerm)

	return &pb.CreateResp{
		Created: true,
	}, nil
}
