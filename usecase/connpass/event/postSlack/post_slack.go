package event

import (
	"context"
	"strconv"
	"time"

	"github.com/ashwanthkumar/slack-go-webhook"
	connpassEvent "github.com/pokotyan/study-slack/infrastructure/connpass/event"
	settingRepo "github.com/pokotyan/study-slack/repository/setting"
	slackUtils "github.com/pokotyan/study-slack/utils/slack"
)

func NewPostSlackImpl(connpassEvent connpassEvent.ConnpassEvent, slackClient *slackUtils.Slack, sr settingRepo.SettingRepository) ConnpassEventUsecase {
	return &connpassEventUsecaseImpl{
		connpassEvent: connpassEvent,
		slackClient:   slackClient,
		settingRepo:   sr,
	}
}

func (u connpassEventUsecaseImpl) PostSlack(ctx context.Context) {
	var param connpassEvent.ReqParam

	s := u.settingRepo.FetchCurrentSetting(ctx)
	isConfigured := u.settingRepo.IsConfigured(s)

	if !isConfigured {
		return
	}

	for i := 0; i < s.SearchRange; i++ {
		day := time.Now()

		t := day.Add(time.Duration(i*24) * time.Hour)
		formatedDate := t.Format("20060102")
		fd, _ := strconv.Atoi(formatedDate)

		param.YmdList = []int{fd}
		param.Keyword = s.Word

		res := u.connpassEvent.Get(param)

		for _, e := range res.Events {
			sendable := e.Accepted >= s.NumOfPeople
			u.slackClient.SendToSlack(e.EventURL, sendable, func(attachment *slack.Attachment) {
				attachment.AddField(slack.Field{Title: "タイトル", Value: e.Title}).AddField(slack.Field{Title: "人数", Value: strconv.Itoa(e.Accepted)}).AddField(slack.Field{Title: "場所", Value: e.Address})
			})
		}
	}
}
