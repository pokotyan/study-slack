package connpassRouter

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/pokotyan/connpass-map-api/middleware/logging"
	"github.com/pokotyan/connpass-map-api/routes/connpass/event"
)

func Init(router *gin.Engine) {
	connpass := router.Group("/connpass")
	{
		connpass.POST("/event", event.Post, middleware.Logging())
	}
}
