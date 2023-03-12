package crypto

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"tools/helpers/cryptoHelper"
	"tools/protos"
	"tools/services/zlog"
)

func GetPwd(c *gin.Context) {
	req := &protos.GetPwdReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		zlog.Error(moduleName, "get pwd", zap.Error(err))
		c.JSON(http.StatusOK, protos.InvalidRequestParameters)
		return
	}
	cryptoHelper.GetPwd(c, req)
}
