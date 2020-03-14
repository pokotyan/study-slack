package event

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	connpassEvent "github.com/pokotyan/study-slack/infrastructure/connpass/event"
	mysql "github.com/pokotyan/study-slack/infrastructure/rdb/client"
	repository "github.com/pokotyan/study-slack/repository/setting"
	usecase "github.com/pokotyan/study-slack/usecase/connpass/event/postSlack"
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
	sr := repository.NewSettingRepository()
	u := usecase.NewPostSlackImpl(ce, sl, sr)

	postSlack(c, u)
}

func postSlack(c *gin.Context, u usecase.ConnpassEventUsecase) {
	db := mysql.Connect()
	defer db.Close()
	db.LogMode(true)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "tx", db)

	u.PostSlack(ctx)

	c.JSON(http.StatusOK, nil)
}
