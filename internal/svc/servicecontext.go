package svc

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shagnyin/gin-scaffold/internal/config"
	"github.com/shagnyin/gin-scaffold/internal/model"
	"github.com/shagnyin/gin-scaffold/pkg/logger"
)

type ServiceContext struct {
	Config config.Config

	DB *model.Queries
}

func NewServiceContext(ctx context.Context, config config.Config) *ServiceContext {
	// 链接数据库
	db, err := sql.Open(config.DB.Driver, config.DB.DataSource)
	if err != nil {
		panic(err)
	}

	// 初始化log
	logger.InitLog(config.Log)

	return &ServiceContext{
		Config: config,

		DB: model.New(db),
	}
}
