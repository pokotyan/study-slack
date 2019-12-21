package status

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	usecase "github.com/pokotyan/connpass-map-api/usecase/status"
)

func Handler(c *gin.Context) {
	apiVer := os.Getenv("API_VERSION")
	status := usecase.GetCurrentStatus()

	c.Writer.Header().Set("X-Api-Revision", apiVer)
	c.JSON(http.StatusOK, status)
}
