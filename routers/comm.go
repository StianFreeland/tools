package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tools/protos"
)

func GetEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, protos.Success())
	})
	setupLogRouter(engine)
	setupIPRouter(engine)
	setupCryptoRouter(engine)
	return engine
}
