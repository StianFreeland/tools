package crypto

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"tools/helpers/cryptoHelper"
	"tools/protos"
	"tools/services/zlog"
)

func GetSHASum(c *gin.Context) {
	req := &protos.SHAEncryptReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		zlog.Error(moduleName, "get sha sum", zap.Error(err))
		c.JSON(http.StatusOK, protos.InvalidRequestParameters)
		return
	}
	req.Msg = strings.TrimSpace(req.Msg)
	cryptoHelper.GetSHASum(c, req)
}
