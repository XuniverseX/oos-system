package main

import (
	"flag"
	"fmt"
	"oos-system/app/rpc/usercenter/internal/utils"
	"oos-system/common/interceptor/rpcserver"

	"oos-system/app/rpc/usercenter/internal/config"
	"oos-system/app/rpc/usercenter/internal/server"
	"oos-system/app/rpc/usercenter/internal/svc"
	"oos-system/app/rpc/usercenter/usercenter"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "app/rpc/usercenter/etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	// 验证码服务重写redis链接
	utils.NewRedisCore(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		usercenter.RegisterUsercenterServer(grpcServer, server.NewUsercenterServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	//rpc log,grpc的全局拦截器
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
