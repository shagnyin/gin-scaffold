package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shagnyin/gin-scaffold/internal/handler/api/demo"
	"github.com/shagnyin/gin-scaffold/internal/svc"
)

func ApiRouter(route *gin.RouterGroup, svcCtx *svc.ServiceContext) {
	v1Router(route.Group("/v1"), svcCtx)
}

func v1Router(route *gin.RouterGroup, svcCtx *svc.ServiceContext) {

	demoGroup := route.Group("demo")
	{
		demoHandler := demo.NewDemo(svcCtx)
		demoGroup.POST("", demoHandler.Create)
		demoGroup.GET("", demoHandler.List)
		demoGroup.DELETE(":id", demoHandler.Delete)
		demoGroup.PUT(":id", demoHandler.Update)
	}
}
