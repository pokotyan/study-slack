package event

import (
	"context"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	connpassEvent "github.com/pokotyan/connpass-map-api/infrastructure/connpass/event"
	usecase "github.com/pokotyan/connpass-map-api/usecase/connpass/event"
)

func PostSlack(c *gin.Context) {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	var res string

	webhookURL := os.Getenv("WEB_HOOK_URL")
	if webhookURL == "" {
		res += "WEB_HOOK_URL is not found."
	}

	searchRange := os.Getenv("SEARCH_RANGE")
	if searchRange == "" {
		res += " SEARCH_RANGE is not found."
	}
	sr, _ := strconv.Atoi(searchRange)

	numOfPeople := os.Getenv("NUM_OF_PEOPLE")
	if numOfPeople == "" {
		res += " NUM_OF_PEOPLE is not found."
	}
	nop, _ := strconv.Atoi(numOfPeople)

	ce := connpassEvent.NewConnpassEvent()
	u := usecase.NewConnpassEventImpl(ce)
	ctx := context.Background()

	if webhookURL != "" && searchRange != "" && numOfPeople != "" {
		u.PostSlack(ctx, webhookURL, nop, sr)
	}

	c.JSON(http.StatusOK, res)
}
