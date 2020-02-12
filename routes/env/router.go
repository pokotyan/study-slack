package envrouter

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/pokotyan/connpass-map-api/middleware/logging"
	"github.com/pokotyan/connpass-map-api/routes/env/set"
)

func Init(router *gin.Engine) {
	connpass := router.Group("/env")
	{
		connpass.POST("", env.Set, middleware.Logging())
	}
}
