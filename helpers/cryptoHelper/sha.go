package cryptoHelper

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
	"tools/protos"
)

func GetSHASum(c *gin.Context, req *protos.SHAEncryptReq) {
	msg := sha256.Sum256([]byte(req.Msg))
	c.JSON(http.StatusOK, protos.Success(&protos.SHAEncryptedInfo{Msg: hex.EncodeToString(msg[:])}))
}
