package logging

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("%v%v\n", c.Request.Host, c.Request.URL)
		c.Next()
	}
}
