package log

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"tools/helpers/logHelper"
	"tools/protos"
	"tools/services/zlog"
)

func GetConfig(c *gin.Context) {
	logHelper.GetConfig(c)
}

func UpdateConfig(c *gin.Context) {
	req := &protos.UpdateLogConfigReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		zlog.Error(moduleName, "update config", zap.Error(err))
		c.JSON(http.StatusOK, protos.InvalidRequestParameters)
		return
	}
	zlog.Warn(moduleName, "update config", zap.Int64("level", req.Level))
	logHelper.UpdateConfig(c, req)
}
