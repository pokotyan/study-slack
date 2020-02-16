package env

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"

	usecase "github.com/pokotyan/connpass-map-api/usecase/connpass/env"
)

type Error struct {
	Name  string `json:"name"`
	Error string `json:"error"`
}
type Empty struct{}

// curl -H "Content-type:application/json" "Accept:application/json" -d '{ "web_hook_url": "https://hooks.slack.com/services/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", "search_range": "2", "num_of_people": "100" }' -X POST http://localhost:7777/env

func Set(c *gin.Context) {
	buf := make([]byte, 2048)
	n, _ := c.Request.Body.Read(buf)
	b := string(buf[0:n])
	str, _ := url.QueryUnescape(b)

	errors := usecase.SetEnv(c, str)

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
