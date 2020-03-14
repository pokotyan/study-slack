package statusrouter

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	usecase "github.com/pokotyan/study-slack/usecase/status"
)

func Handler(c *gin.Context) {
	apiVer := os.Getenv("API_VERSION")
	status := usecase.GetCurrentStatus()

	c.Writer.Header().Set("X-Api-Revision", apiVer)
	c.JSON(http.StatusOK, status)
}
