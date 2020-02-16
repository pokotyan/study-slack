package connpass

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	connpassEvent "github.com/pokotyan/connpass-map-api/infrastructure/connpass/event"
	usecase "github.com/pokotyan/connpass-map-api/usecase/connpass/event"
	slackUtils "github.com/pokotyan/connpass-map-api/utils/slack"
)

func postSlack(cmd *cobra.Command, args []string) error {
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
	sl, _ := slackUtils.NewSlack(webhookURL)
	u := usecase.NewPostSlackImpl(ce, sl)
	ctx := context.Background()

	if webhookURL != "" && searchRange != "" && numOfPeople != "" {
		u.PostSlack(ctx, nop, sr)
	}

	fmt.Printf("%v", res)

	return nil
}
