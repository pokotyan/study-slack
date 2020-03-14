package env

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nlopes/slack"

	mysql "github.com/pokotyan/study-slack/infrastructure/rdb/client"
	repository "github.com/pokotyan/study-slack/repository/setting"
	usecase "github.com/pokotyan/study-slack/usecase/connpass/env/openDialog"
)

type Payload struct {
	Token       string `form:"token"`
	TeamID      string `form:"team_id"`
	TeamDomain  string `form:"team_domain"`
	ChannelID   string `form:"channel_id"`
	ChannelName string `form:"channel_name"`
	UserID      string `form:"user_id"`
	UserName    string `form:"user_name"`
	Command     string `form:"command"`
	Text        string `form:"text"`
	ResponseURL string `form:"response_url"`
	TriggerID   string `form:"trigger_id"`
}

func OpenDialog(c *gin.Context) {
	slackToken := os.Getenv("SLACK_TOKEN")
	if slackToken == "" {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	slackClient := slack.New(slackToken)
	sr := repository.NewSettingRepository()
	u := usecase.NewOpenDialogImpl(slackClient, sr)

	openDialog(c, u)
}

func openDialog(c *gin.Context, u usecase.ConnpassEnvUsecase) {
	db := mysql.Connect()
	defer db.Close()
	db.LogMode(true)

	var body Payload
	c.Bind(&body)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "tx", db)

	err := u.OpenDialog(ctx, body.UserID, body.TriggerID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	return
}
