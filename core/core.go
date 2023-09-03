package main

import (
	"flag"
	"fmt"

	"NetDisk/core/internal/config"
	"NetDisk/core/internal/handler"
	"NetDisk/core/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/core-api.yaml", "the config file")

func main() {
	// @title ToDoList API
	// @version 0.0.1
	// @description This is a sample Server pets
	// @securityDefinitions.apikey ApiKeyAuth
	// @name FanOne
	// @BasePath /api/v1
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
