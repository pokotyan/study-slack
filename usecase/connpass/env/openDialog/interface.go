package env

import (
	"context"

	"github.com/nlopes/slack"

	settingRepository "github.com/pokotyan/study-slack/repository/setting"
)

type (
	ConnpassEnvUsecase interface {
		OpenDialog(ctx context.Context, userID string, triggerID string) error
	}
	connpassEnvUsecaseImpl struct {
		slackClient *slack.Client
		settingRepo settingRepository.SettingRepository
	}
)
