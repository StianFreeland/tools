package cryptoHelper

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"tools/protos"
	"tools/services/zlog"
)

func GetKey(c *gin.Context, req *protos.GetKeyReq) {
	key, err := rsa.GenerateKey(rand.Reader, req.Length)
	if err != nil {
		zlog.Error(moduleName, "get key", zap.Error(err))
		c.JSON(http.StatusOK, protos.InvalidKeyLength)
	}
	mPvtKey := x509.MarshalPKCS1PrivateKey(key)
	pvtKey := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: mPvtKey,
	})
	mPubKey, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		zlog.Error(moduleName, "get key", zap.Error(err))
		c.JSON(http.StatusOK, protos.InternalServerError)
	}
	pubKey := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: mPubKey,
	})
	c.JSON(http.StatusOK, protos.Success(&protos.KeyInfo{PvtKey: string(pvtKey), PubKey: string(pubKey)}))
}
