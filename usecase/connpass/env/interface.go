package env

import (
	"context"

	"github.com/nlopes/slack"
)

type (
	ConnpassEnvUsecase interface {
		OpenDialog(ctx context.Context, userID string, triggerID string) error
		// SetEnv(ctx context.Context, rawBody string) []EnvError
	}
	connpassEnvUsecaseImpl struct {
		slackClient *slack.Client
	}
)
