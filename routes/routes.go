package routes

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/pokotyan/connpass-map-api/middleware/logging"
	connpassrouter "github.com/pokotyan/connpass-map-api/routes/connpass"
	envrouter "github.com/pokotyan/connpass-map-api/routes/env"
	statusrouter "github.com/pokotyan/connpass-map-api/routes/status"
)

func Init(router *gin.Engine) {
	router.GET("/status", statusrouter.Handler, middleware.Logging())
	connpassrouter.Init(router)
	envrouter.Init(router)
}
