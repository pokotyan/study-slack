package event

import (
	"context"
	"strconv"
	"time"

	"github.com/ashwanthkumar/slack-go-webhook"
	connpassEvent "github.com/pokotyan/study-slack/infrastructure/connpass/event"
	slackUtils "github.com/pokotyan/study-slack/utils/slack"
)

func NewPostSlackImpl(connpassEvent connpassEvent.ConnpassEvent, slackClient *slackUtils.Slack) ConnpassEventUsecase {
	return &connpassEventUsecaseImpl{
		connpassEvent: connpassEvent,
		slackClient:   slackClient,
	}
}

func (u connpassEventUsecaseImpl) PostSlack(ctx context.Context, nop int, searchRange int) {
	var param connpassEvent.ReqParam

	for i := 0; i < searchRange; i++ {
		day := time.Now()

		t := day.Add(time.Duration(i*24) * time.Hour)
		formatedDate := t.Format("20060102")
		fd, _ := strconv.Atoi(formatedDate)

		param.YmdList = []int{fd}

		res := u.connpassEvent.Get(param)

		for _, e := range res.Events {
			sendable := e.Accepted >= nop
			u.slackClient.SendToSlack(e.EventURL, sendable, func(attachment *slack.Attachment) {
				attachment.AddField(slack.Field{Title: "タイトル", Value: e.Title}).AddField(slack.Field{Title: "人数", Value: strconv.Itoa(e.Accepted)}).AddField(slack.Field{Title: "場所", Value: e.Address})
			})
		}
	}
}
