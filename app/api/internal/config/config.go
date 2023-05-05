package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	UsercenterRpc  zrpc.RpcClientConf
	BucketRpc      zrpc.RpcClientConf
	FileserviceRpc zrpc.RpcClientConf
	MetadataRpc    zrpc.RpcClientConf
	TempPath       string
}
