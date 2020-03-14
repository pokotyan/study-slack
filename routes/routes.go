package routes

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/pokotyan/study-slack/middleware/logging"
	connpassrouter "github.com/pokotyan/study-slack/routes/connpass"
	statusrouter "github.com/pokotyan/study-slack/routes/status"
)

func Init(router *gin.Engine) {
	router.GET("/status", statusrouter.Handler, middleware.Logging())
	connpassrouter.Init(router)
}
