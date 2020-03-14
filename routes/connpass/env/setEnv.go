package env

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	mysql "github.com/pokotyan/study-slack/infrastructure/rdb/client"
	repository "github.com/pokotyan/study-slack/repository/setting"
	usecase "github.com/pokotyan/study-slack/usecase/connpass/env/setEnv"
	gu "github.com/pokotyan/study-slack/utils/gin"
)

type Error struct {
	Name  string `json:"name"`
	Error string `json:"error"`
}
type Empty struct{}

// curl -H "Content-type:application/json" "Accept:application/json" -d '{ "web_hook_url": "https://hooks.slack.com/services/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", "search_range": "2", "num_of_people": "100" }' -X POST http://localhost:7777/env

func Set(c *gin.Context) {
	sr := repository.NewSettingRepository()
	u := usecase.NewSetEnvImpl(sr)

	set(c, u)
}

func set(c *gin.Context, u usecase.ConnpassEnvUsecase) {
	db := mysql.Connect()
	defer db.Close()
	db.LogMode(true)

	str := gu.GetRawBody(c)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "tx", db)
	errors := u.SetEnv(ctx, str)

	validationErrors := []Error{}

	for _, err := range errors {
		errors := append(validationErrors, Error{Name: err.ID, Error: err.Msg})
		validationErrors = errors
	}

	if len(validationErrors) != 0 {
		c.JSON(http.StatusOK, gin.H{"errors": validationErrors})
		return
	}

	c.JSON(http.StatusOK, Empty{})
}
