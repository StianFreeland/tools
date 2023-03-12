package logHelper

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tools/protos"
	"tools/services/zlog"
)

func UpdateConfig(c *gin.Context, req *protos.UpdateLogConfigReq) {
	zlog.UpdateLevel(req.Level)
	c.JSON(http.StatusOK, protos.Success())
}
