package env

import (
	"context"

	"github.com/nlopes/slack"
)

type (
	ConnpassEnvUsecase interface {
		OpenDialog(ctx context.Context, userID string, triggerID string) error
	}
	connpassEnvUsecaseImpl struct {
		slackClient *slack.Client
	}
)
