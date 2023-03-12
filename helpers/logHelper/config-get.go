package logHelper

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tools/protos"
	"tools/services/zlog"
)

func GetConfig(c *gin.Context) {
	data := protos.LogConfig{
		Level: zlog.GetLevel(),
	}
	c.JSON(http.StatusOK, protos.Success(data))
}
