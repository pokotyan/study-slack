package event

import (
	"context"

	connpassEvent "github.com/pokotyan/connpass-map-api/infrastructure/connpass/event"
	slackUtils "github.com/pokotyan/connpass-map-api/utils/slack"
)

type (
	ConnpassEventUsecase interface {
		GetEvent(ctx context.Context, param connpassEvent.ReqParam) connpassEvent.Res
		PostSlack(ctx context.Context, nop int, searchRange int)
	}
	connpassEventUsecaseImpl struct {
		connpassEvent connpassEvent.ConnpassEvent
		slackClient   *slackUtils.Slack
	}
)
