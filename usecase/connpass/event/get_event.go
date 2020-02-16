package event

import (
	"context"

	connpassEvent "github.com/pokotyan/connpass-map-api/infrastructure/connpass/event"
)

func NewGetEventImpl(connpassEvent connpassEvent.ConnpassEvent) ConnpassEventUsecase {
	return &connpassEventUsecaseImpl{
		connpassEvent: connpassEvent,
	}
}

func (u connpassEventUsecaseImpl) GetEvent(ctx context.Context, param connpassEvent.ReqParam) connpassEvent.Res {
	// @todo ctxにdbを含める

	res := u.connpassEvent.Get(param)

	return res
}
