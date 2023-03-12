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

func GetMD5Sum(c *gin.Context) {
	req := &protos.MD5EncryptReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		zlog.Error(moduleName, "get md5 sum", zap.Error(err))
		c.JSON(http.StatusOK, protos.InvalidRequestParameters)
		return
	}
	req.Msg = strings.TrimSpace(req.Msg)
	cryptoHelper.GetMD5Sum(c, req)
}
