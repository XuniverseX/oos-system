package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"oos-system/app/rpc/metadata/internal/config"
	"oos-system/app/rpc/metadata/internal/server"
	"oos-system/app/rpc/metadata/internal/svc"
	"oos-system/app/rpc/metadata/pb"
	"oos-system/common/interceptor/rpcserver"
)

var configFile = flag.String("f", "app/rpc/metadata/etc/metadata.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterMetadataServiceServer(grpcServer, server.NewMetadataServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
