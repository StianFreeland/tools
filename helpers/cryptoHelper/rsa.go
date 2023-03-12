package cryptoHelper

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"tools/comm"
	"tools/protos"
	"tools/services/zlog"
)

func RSAEncrypt(c *gin.Context, req *protos.RSAEncryptReq) {
	msg, err := comm.RSAEncryptMsg(req.Msg, []byte(req.Key))
	if err != nil {
		zlog.Error(moduleName, "rsa encrypt", zap.Error(err))
		c.JSON(http.StatusOK, protos.Error(err))
		return
	}
	c.JSON(http.StatusOK, protos.Success(&protos.RSAEncryptedInfo{Msg: msg}))
}

func RSADecrypt(c *gin.Context, req *protos.RSADecryptReq) {
	msg, err := comm.RSADecryptMsg(req.Msg, []byte(req.Key))
	if err != nil {
		zlog.Error(moduleName, "rsa decrypt", zap.Error(err))
		c.JSON(http.StatusOK, protos.Error(err))
		return
	}
	c.JSON(http.StatusOK, protos.Success(&protos.RSADecryptedInfo{Msg: msg}))
}
