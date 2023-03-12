package cryptoHelper

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"tools/protos"
)

const allCharSet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GetPwd(c *gin.Context, req *protos.GetPwdReq) {
	pwd := ""
	for i := 0; i < req.Length; i++ {
		pwd += string(allCharSet[rand.Intn(len(allCharSet))])
	}
	c.JSON(http.StatusOK, protos.Success(&protos.PwdInfo{Pwd: pwd}))
}
