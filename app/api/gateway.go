package main

import (
	"flag"
	"fmt"

	"oos-system/app/api/internal/config"
	"oos-system/app/api/internal/handler"
	"oos-system/app/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "app/api/etc/gateway.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 跨域处理 之后将后边的域名替换
	server := rest.MustNewServer(c.RestConf, rest.WithCors("*"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
