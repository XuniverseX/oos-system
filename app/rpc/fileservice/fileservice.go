package main

import (
	"flag"
	"fmt"
	"oos-system/common/interceptor/rpcserver"

	"oos-system/app/rpc/fileservice/internal/config"
	"oos-system/app/rpc/fileservice/internal/server"
	"oos-system/app/rpc/fileservice/internal/svc"
	"oos-system/app/rpc/fileservice/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "app/rpc/fileservice/etc/fileservice.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	config.Init(c)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterFileserviceServer(grpcServer, server.NewFileserviceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	s.AddOptions(grpc.MaxRecvMsgSize(10485760))
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
