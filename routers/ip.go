package routers

import (
	"github.com/gin-gonic/gin"
	"tools/handlers/ip"
)

func setupIPRouter(engine *gin.Engine) {
	engine.GET("/ip", ip.GetIPInfo)
}
