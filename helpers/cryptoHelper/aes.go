package cryptoHelper

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"tools/comm"
	"tools/protos"
	"tools/services/zlog"
)

func AESEncrypt(c *gin.Context, req *protos.AESEncryptReq) {
	msg, err := comm.AESEncryptMsg(req.Msg, []byte(req.Key), []byte(req.IV))
	if err != nil {
		zlog.Error(moduleName, "aes encrypt", zap.Error(err))
		c.JSON(http.StatusOK, protos.Error(err))
		return
	}
	c.JSON(http.StatusOK, protos.Success(&protos.AESEncryptedInfo{Msg: msg}))
}

func AESDecrypt(c *gin.Context, req *protos.AESDecryptReq) {
	msg, err := comm.AESDecryptMsg(req.Msg, []byte(req.Key), []byte(req.IV))
	if err != nil {
		zlog.Error(moduleName, "aes decrypt", zap.Error(err))
		c.JSON(http.StatusOK, protos.Error(err))
		return
	}
	c.JSON(http.StatusOK, protos.Success(&protos.AESDecryptedInfo{Msg: msg}))
}
