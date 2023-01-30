package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/shagnyin/gin-scaffold/internal"
	"github.com/shagnyin/gin-scaffold/internal/config"
	"github.com/shagnyin/gin-scaffold/internal/svc"
	"github.com/shagnyin/gin-scaffold/pkg/conf"
	"log"
)

var configFile = flag.String("f", "config/config.yaml", "the conf file")

func main() {
	flag.Parse()

	// 加载配置项
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 初始化项目
	var ctx context.Context
	serviceContext := svc.NewServiceContext(ctx, c)

	// 设置环境
	gin.SetMode(c.Mode)

	// 初始化gin
	engine := gin.Default()

	// 引入路由文件
	internal.Router(engine, serviceContext)

	// 启动项目
	log.Fatal(engine.Run(c.Host + ":" + c.Port))
}
