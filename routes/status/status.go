package status

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecase "github.com/pokotyan/connpass-map-api/usecase/status"
)

func Handler(c *gin.Context) {
	status := usecase.GetCurrentStatus()

	c.JSON(http.StatusOK, status)
}
