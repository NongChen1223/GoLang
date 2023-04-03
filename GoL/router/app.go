package router

import (
	"ginchat/docs"
	"ginchat/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	// 1.创建路由
	r := gin.Default()
	//swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/index", service.GetIndex)
	r.GET("/user/GetUserList", service.GetUserList)
	return r
}
