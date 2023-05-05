package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	TempPath string
	//WriteToPath string
	MaxThread int64
}
