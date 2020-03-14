package event

import (
	"context"

	connpassEvent "github.com/pokotyan/study-slack/infrastructure/connpass/event"
	settingRepository "github.com/pokotyan/study-slack/repository/setting"
	slackUtils "github.com/pokotyan/study-slack/utils/slack"
)

type (
	ConnpassEventUsecase interface {
		PostSlack(ctx context.Context)
	}
	connpassEventUsecaseImpl struct {
		connpassEvent connpassEvent.ConnpassEvent
		slackClient   *slackUtils.Slack
		settingRepo   settingRepository.SettingRepository
	}
)
