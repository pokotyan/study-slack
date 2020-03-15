package connpass

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	connpassEvent "github.com/pokotyan/study-slack/infrastructure/connpass/event"
	mysql "github.com/pokotyan/study-slack/infrastructure/rdb/client"
	repository "github.com/pokotyan/study-slack/repository/setting"
	usecase "github.com/pokotyan/study-slack/usecase/connpass/event/postSlack"
	slackUtils "github.com/pokotyan/study-slack/utils/slack"
)

func PostSlack(cmd *cobra.Command, args []string) error {
	webhookURL := os.Getenv("WEB_HOOK_URL")
	if webhookURL == "" {
		fmt.Println("WEB_HOOK_URL is not found.")
		return nil
	}

	ce := connpassEvent.NewConnpassEvent()
	sl, _ := slackUtils.NewSlack(webhookURL)
	sr := repository.NewSettingRepository()
	u := usecase.NewPostSlackImpl(ce, sl, sr)

	return postSlack(u)
}

func postSlack(u usecase.ConnpassEventUsecase) error {
	db := mysql.Connect()
	defer db.Close()
	db.LogMode(true)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "tx", db)

	u.PostSlack(ctx)

	return nil
}
