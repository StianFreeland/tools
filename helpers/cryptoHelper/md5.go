package cryptoHelper

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
	"tools/protos"
)

func GetMD5Sum(c *gin.Context, req *protos.MD5EncryptReq) {
	msg := md5.Sum([]byte(req.Msg))
	c.JSON(http.StatusOK, protos.Success(&protos.MD5EncryptedInfo{Msg: hex.EncodeToString(msg[:])}))
}
