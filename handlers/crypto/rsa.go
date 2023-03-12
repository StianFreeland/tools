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

func RSAEncrypt(c *gin.Context) {
	req := &protos.RSAEncryptReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		zlog.Error(moduleName, "rsa encrypt", zap.Error(err))
		c.JSON(http.StatusOK, protos.InvalidRequestParameters)
		return
	}
	req.Msg = strings.TrimSpace(req.Msg)
	req.Key = strings.TrimSpace(req.Key)
	cryptoHelper.RSAEncrypt(c, req)
}

func RSADecrypt(c *gin.Context) {
	req := &protos.RSADecryptReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		zlog.Error(moduleName, "rsa decrypt", zap.Error(err))
		c.JSON(http.StatusOK, protos.InvalidRequestParameters)
		return
	}
	req.Msg = strings.TrimSpace(req.Msg)
	req.Key = strings.TrimSpace(req.Key)
	cryptoHelper.RSADecrypt(c, req)
}
