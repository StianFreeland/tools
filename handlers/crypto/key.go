package crypto

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"tools/helpers/cryptoHelper"
	"tools/protos"
	"tools/services/zlog"
)

func GetKey(c *gin.Context) {
	req := &protos.GetKeyReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		zlog.Error(moduleName, "get key", zap.Error(err))
		c.JSON(http.StatusOK, protos.InvalidRequestParameters)
		return
	}
	cryptoHelper.GetKey(c, req)
}
