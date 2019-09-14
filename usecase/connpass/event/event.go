package event

import (
	conpassEvent "github.com/pokotyan/connpass-map-api/infrastructure/connpass/event"
)

func GetEvent(param conpassEvent.ReqParam) conpassEvent.Res {
	res := conpassEvent.Get(param)

	return res
}
