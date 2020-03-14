package server

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	routes "github.com/pokotyan/study-slack/routes"
)

func Init() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	routes.Init(router)

	router.Run(":" + os.Getenv("PORT"))
}
