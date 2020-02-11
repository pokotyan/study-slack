package routes

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/pokotyan/connpass-map-api/middleware/logging"
	connpassRouter "github.com/pokotyan/connpass-map-api/routes/connpass"
	envRouter "github.com/pokotyan/connpass-map-api/routes/env"
	"github.com/pokotyan/connpass-map-api/routes/status"
)

func Init(router *gin.Engine) {
	router.GET("/status", status.Handler, middleware.Logging())
	connpassRouter.Init(router)
	envRouter.Init(router)
}
