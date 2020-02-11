package event

import (
	"context"
	"strconv"
	"time"

	connpassEvent "github.com/pokotyan/connpass-map-api/infrastructure/connpass/event"
	"github.com/pokotyan/connpass-map-api/utils"
)

func NewConnpassEventImpl(connpassEvent connpassEvent.ConnpassEvent) ConnpassEventUsecase {
	return &connpassEventUsecaseImpl{
		connpassEvent: connpassEvent,
	}
}

func (u connpassEventUsecaseImpl) GetEvent(ctx context.Context, param connpassEvent.ReqParam) connpassEvent.Res {
	// @todo ctxにdbを含める

	res := u.connpassEvent.Get(param)

	return res
}

func (u connpassEventUsecaseImpl) PostSlack(ctx context.Context, webhookURL string, nop int, searchRange int) {
	var param connpassEvent.ReqParam

	for i := 0; i < searchRange; i++ {
		day := time.Now()

		t := day.Add(time.Duration(i*24) * time.Hour)
		formatedDate := t.Format("20060102")
		fd, _ := strconv.Atoi(formatedDate)

		param.YmdList = []int{fd}

		res := u.GetEvent(ctx, param)

		for _, e := range res.Events {
			sl, _ := utils.NewSlack(webhookURL)

			sl.SendToSlack(e.EventURL, func() bool { return e.Accepted >= nop }())
		}
	}
}
