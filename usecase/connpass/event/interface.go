package event

import (
	"context"

	connpassEvent "github.com/pokotyan/connpass-map-api/infrastructure/connpass/event"
)

type (
	ConnpassEventUsecase interface {
		GetEvent(ctx context.Context, param connpassEvent.ReqParam) connpassEvent.Res
		PostSlack(ctx context.Context, webhookURL string, nop int, searchRange int)
	}
	connpassEventUsecaseImpl struct {
		connpassEvent connpassEvent.ConnpassEvent
	}
)
