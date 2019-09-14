package server

import (
	"github.com/gin-gonic/gin"
	routes "github.com/pokotyan/connpass-map-api/routes"
)

func Init() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.Init(router)

	router.Run(":7777")
}
