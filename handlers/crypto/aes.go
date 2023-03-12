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

func AESEncrypt(c *gin.Context) {
	req := &protos.AESEncryptReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		zlog.Error(moduleName, "aes encrypt", zap.Error(err))
		c.JSON(http.StatusOK, protos.InvalidRequestParameters)
		return
	}
	req.Msg = strings.TrimSpace(req.Msg)
	req.Key = strings.TrimSpace(req.Key)
	req.IV = strings.TrimSpace(req.IV)
	cryptoHelper.AESEncrypt(c, req)
}

func AESDecrypt(c *gin.Context) {
	req := &protos.AESDecryptReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		zlog.Error(moduleName, "aes decrypt", zap.Error(err))
		c.JSON(http.StatusOK, protos.InvalidRequestParameters)
		return
	}
	req.Msg = strings.TrimSpace(req.Msg)
	req.Key = strings.TrimSpace(req.Key)
	req.IV = strings.TrimSpace(req.IV)
	zlog.Info(
		moduleName, "aes decrypt",
		zap.String("msg", req.Msg),
		zap.String("key", req.Key),
		zap.String("iv", req.IV),
	)
	cryptoHelper.AESDecrypt(c, req)
}
