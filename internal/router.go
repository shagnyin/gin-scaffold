package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/shagnyin/gin-scaffold/internal/middleware"
	"github.com/shagnyin/gin-scaffold/internal/router"
	"github.com/shagnyin/gin-scaffold/internal/svc"
)

func Router(engine *gin.Engine, svcCtx *svc.ServiceContext) {
	engine.Use(middleware.CORSMiddleware(), middleware.TraceMiddleware())

	// api router group
	router.ApiRouter(engine.Group("api"), svcCtx)

	//engine.NoRoute(func(context *gin.Context) {
	//	context.String(http.StatusNotFound, "you miss load")
	//})
}
