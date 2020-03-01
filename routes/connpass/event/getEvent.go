package event

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	connpassEvent "github.com/pokotyan/study-slack/infrastructure/connpass/event"
	usecase "github.com/pokotyan/study-slack/usecase/connpass/event"
)

// curl -H "Content-type:application/json" "Accept:application/json" -d '{ "Keyword": "python", "YmList": [201209] }' -X POST http://localhost:7777/connpass/event | jq .

type Body struct {
	Keyword string `json:"keyword"`
	YmList  []int  `json:"ymList"`
	YmdList []int  `json:"ymdList"`
}

func GetEvent(c *gin.Context) {
	ce := connpassEvent.NewConnpassEvent()
	u := usecase.NewGetEventImpl(ce)

	getEvent(c, u)
}

func getEvent(c *gin.Context, u usecase.ConnpassEventUsecase) {
	var body Body
	c.BindJSON(&body)

	var param connpassEvent.ReqParam
	param.Keyword = body.Keyword
	param.YmList = body.YmList
	param.YmdList = body.YmdList

	ctx := context.Background()
	res := u.GetEvent(ctx, param)

	c.JSON(http.StatusOK, res)
}
