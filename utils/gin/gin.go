package gu

import (
	"net/url"

	"github.com/gin-gonic/gin"
)

func GetRawBody(c *gin.Context) string {
	buf := make([]byte, 2048)
	n, _ := c.Request.Body.Read(buf)
	b := string(buf[0:n])
	str, _ := url.QueryUnescape(b)

	return str
}
