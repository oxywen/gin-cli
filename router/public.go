package router

import (
	v1 "gin-server-cli/api/v1"
	"github.com/gin-gonic/gin"
)

func InitPublicRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	//公开的，不做鉴权的路由
	Router.GET("/login", v1.Login)
	return Router
}
