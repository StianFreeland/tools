package routers

import (
	"github.com/gin-gonic/gin"
	"tools/handlers/log"
)

func setupLogRouter(engine *gin.Engine) {
	engine.GET("/log/config", log.GetConfig)
	engine.PUT("/log/config", log.UpdateConfig)
}
