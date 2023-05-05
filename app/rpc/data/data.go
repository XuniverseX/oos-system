package main

import (
	"flag"
	"fmt"
	"os"

	"oos-system/app/rpc/data/internal/config"
	"oos-system/app/rpc/data/internal/server"
	"oos-system/app/rpc/data/internal/svc"
	"oos-system/app/rpc/data/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/data.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	if _, err := os.Stat(c.RootPath); os.IsNotExist(err) {
		os.MkdirAll(c.RootPath, os.ModePerm)
	}

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterDataServer(grpcServer, server.NewDataServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
