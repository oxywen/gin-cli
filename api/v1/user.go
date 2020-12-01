package v1

import (
	"gin-server-cli/core/constant"
	"gin-server-cli/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(context *gin.Context) {
	name := context.Param("name")
	msg := service.Hello(name)
	context.JSON(http.StatusOK, gin.H{
		"code": constant.SUCCESS,
		"msg":  msg,
	})
}
