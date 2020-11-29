package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(context *gin.Context) {
	context.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"msg":  "login success",
		"data": nil,
	})
}
