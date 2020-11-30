package v1

import (
	"gin-server-cli/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Login(context *gin.Context) {
	id, _ := strconv.ParseInt(context.Param("id"), 10, 64)
	community, _ := service.GetCommunityById(id)
	context.JSON(http.StatusOK, community)
}
