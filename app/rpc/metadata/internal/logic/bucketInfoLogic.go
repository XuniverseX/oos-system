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
	"path/filepath"
	"strings"
	"sync"
)

type BucketInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBucketInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BucketInfoLogic {
	return &BucketInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BucketInfoLogic) BucketInfo(in *pb.BucketInfoReq) (*pb.BucketInfoResp, error) {
	//node := &pb.ObjectInfo{ObjectName: in.BucketName + in.ObjectName}
	location := path.Join([]string{l.svcCtx.Config.BaseLocation, in.BucketName, in.ObjectName}...)
	_, err := os.Stat(location)
	var children []*pb.ObjectInfo
	if err == nil {
		children = scanDir(location, l.svcCtx.RWLockMap)
	} else if !os.IsNotExist(err) {
		return nil, errors.Wrapf(xerr.NewErrMsg("信息不存在"), "bucketName is %s", in.BucketName+in.ObjectName)
	}

	Sort(children)
	return &pb.BucketInfoResp{
		ObjectInfo: children,
	}, nil
}

func scanDir(path string, maps *sync.Map) []*pb.ObjectInfo {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
		return nil
	}
	var children []*pb.ObjectInfo
	for _, e := range dir {
		if e.IsDir() {
			// 拿到桶或者对象对应的读写锁， 不存在就创建
			var rwLock sync.RWMutex
			var lock = &rwLock
			if mutex, ok := maps.LoadOrStore(path+e.Name(), &rwLock); ok {
				lock = mutex.(*sync.RWMutex)
			}
			treeNode := &pb.ObjectInfo{}
			lock.RLock()
			readFile, _ := ioutil.ReadFile(filepath.Join(path, e.Name(), constent.MetaFileName))
			lock.RUnlock()
			objectInfo := &metaInfo.ObjectInfo{}
			jsonx.Unmarshal(readFile, objectInfo)
			treeNode.Size = objectInfo.Size
			treeNode.CreateTime = objectInfo.CreateTime
			treeNode.ObjectName = objectInfo.ObjectName
			treeNode.Name = objectInfo.ObjectName
			treeNode.IsDir = objectInfo.IsDir
			treeNode.Hash = objectInfo.ChekCode
			spilts := strings.Split(objectInfo.ObjectName, "/")
			treeNode.Name = spilts[len(spilts)-1]
			children = append(children, treeNode)
		}
	}
	return children
}

func Sort(arr []*pb.ObjectInfo) {
	n := len(arr)

	slow := 0
	fast := 0

	for ; fast < n; fast++ {
		if arr[fast].IsDir {
			tmp := arr[slow]
			arr[slow] = arr[fast]
			arr[fast] = tmp
			slow++
		}
	}
}
