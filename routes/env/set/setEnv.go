package env

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Body struct {
	WebhookURL  string `json:"web_hook_url"`
	SearchRange string `json:"search_range"`
	NumOfPeople string `json:"num_of_people"`
}

// curl -H "Content-type:application/json" "Accept:application/json" -d '{ "web_hook_url": "https://hooks.slack.com/services/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", "search_range": "2", "num_of_people": "100" }' -X POST http://localhost:7777/env

// @todo slackのダイアログで設定値を入力し、このapiを叩くことで環境変数を更新する

func Set(c *gin.Context) {
	var body Body
	c.BindJSON(&body)

	os.Setenv("WEB_HOOK_URL", body.WebhookURL)
	os.Setenv("SEARCH_RANGE", body.SearchRange)
	os.Setenv("NUM_OF_PEOPLE", body.NumOfPeople)

	c.JSON(http.StatusOK, nil)
}
