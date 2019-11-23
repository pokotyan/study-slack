package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	routes "github.com/pokotyan/connpass-map-api/routes"
)

func Init() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	routes.Init(router)

	router.Run(":7777")
}
