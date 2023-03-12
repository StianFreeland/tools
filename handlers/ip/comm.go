package ip

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tools/protos"
)

func GetIPInfo(c *gin.Context) {
	data := &protos.IPInfo{Address: c.ClientIP()}
	c.JSON(http.StatusOK, protos.Success(data))
}
