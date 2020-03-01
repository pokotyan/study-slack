package env

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nlopes/slack"

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

	var body Payload
	c.Bind(&body)

	slackClient := slack.New(slackToken)
	u := usecase.NewOpenDialogImpl(slackClient)
	err := u.OpenDialog(c, body.UserID, body.TriggerID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	return
}
