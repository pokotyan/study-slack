package event

import (
	"context"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	connpassEvent "github.com/pokotyan/study-slack/infrastructure/connpass/event"
	usecase "github.com/pokotyan/study-slack/usecase/connpass/event"
	slackUtils "github.com/pokotyan/study-slack/utils/slack"
)

// curl -H "Content-type:application/json" "Accept:application/json" -d '' -X POST http://localhost:7777/connpass/slack

func PostSlack(c *gin.Context) {
	webhookURL := os.Getenv("WEB_HOOK_URL")
	if webhookURL == "" {
		c.JSON(http.StatusOK, "WEB_HOOK_URL is not found.")

		return
	}

	ce := connpassEvent.NewConnpassEvent()
	sl, _ := slackUtils.NewSlack(webhookURL)
	u := usecase.NewPostSlackImpl(ce, sl)

	postSlack(c, u)
}

func postSlack(c *gin.Context, u usecase.ConnpassEventUsecase) {
	var res string

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

	ctx := context.Background()

	if searchRange != "" && numOfPeople != "" {
		u.PostSlack(ctx, nop, sr)
	}

	c.JSON(http.StatusOK, res)
}
