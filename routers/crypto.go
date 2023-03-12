package routers

import (
	"github.com/gin-gonic/gin"
	"tools/handlers/crypto"
)

func setupCryptoRouter(engine *gin.Engine) {
	engine.GET("/crypto/pwd", crypto.GetPwd)
	engine.GET("/crypto/key", crypto.GetKey)
	engine.GET("/crypto/md5", crypto.GetMD5Sum)
	engine.GET("/crypto/sha", crypto.GetSHASum)
	engine.GET("/crypto/aes-encrypt", crypto.AESEncrypt)
	engine.GET("/crypto/aes-decrypt", crypto.AESDecrypt)
	engine.GET("/crypto/rsa-encrypt", crypto.RSAEncrypt)
	engine.GET("/crypto/rsa-decrypt", crypto.RSADecrypt)
}
