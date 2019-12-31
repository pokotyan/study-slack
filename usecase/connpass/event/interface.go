package event

import (
	"context"

	connpassEvent "github.com/pokotyan/connpass-map-api/infrastructure/connpass/event"
)

type (
	ConnpassEventUsecase interface {
		GetEvent(ctx context.Context, param connpassEvent.ReqParam) connpassEvent.Res
	}
	connpassEventUsecaseImpl struct {
		connpassEvent connpassEvent.ConnpassEvent
	}
)
