package event

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	connpassEvent "github.com/pokotyan/connpass-map-api/infrastructure/connpass/event"
	usecase "github.com/pokotyan/connpass-map-api/usecase/connpass/event"
)

// curl -H "Content-type:application/json" "Accept:application/json" -d '{ "Keyword": "python", "YmList": [201209] }' -X POST http://localhost:7777/connpass/event | jq .

type Body struct {
	Keyword string `json:keyword`
	YmList  []int  `json:ymList`
	YmdList []int  `json:ymdList`
}

func Post(c *gin.Context) {
	var body Body
	c.BindJSON(&body)

	var param connpassEvent.ReqParam
	param.Keyword = body.Keyword
	param.YmList = body.YmList
	param.YmdList = body.YmdList

	u := usecase.AssignGetEventWithUsecase()
	ctx := context.Background()
	// ctx = context.WithValue(ctx, "tx", db)
	res := u.GetEvent(ctx, param)

	c.JSON(http.StatusOK, res)
}
