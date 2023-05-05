package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"oos-system/app/api/internal/config"
	"oos-system/app/rpc/bucket/bucketclient"
	"oos-system/app/rpc/fileservice/fileservice"
	"oos-system/app/rpc/metadata/metadataservice"
	"oos-system/app/rpc/usercenter/usercenterclient"
	"time"
)

type ServiceContext struct {
	Config         config.Config
	UsercenterRpc  usercenterclient.Usercenter
	BucketRpc      bucketclient.Bucket
	FileserviceRpc fileservice.Fileservice
	MetadataRpc    metadataservice.MetadataService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UsercenterRpc:  usercenterclient.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpc)),
		BucketRpc:      bucketclient.NewBucket(zrpc.MustNewClient(c.BucketRpc)),
		FileserviceRpc: fileservice.NewFileservice(zrpc.MustNewClient(c.FileserviceRpc, zrpc.WithTimeout(time.Second*10))),
		MetadataRpc:    metadataservice.NewMetadataService(zrpc.MustNewClient(c.MetadataRpc)),
	}
}
