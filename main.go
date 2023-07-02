package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yog-singh/gandharva/src/config"
	"github.com/yog-singh/gandharva/src/db"
	"github.com/yog-singh/gandharva/src/router"
	"github.com/yog-singh/gandharva/src/services"
)

var (
	server *gin.Engine
)

func init() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	switch config.Environment {
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	}

	db.ConnectDB(&config)

	server = gin.Default()
	services.RunScheduledHeartbeatCheckService()
}

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	router.InitRoutes(server)
	log.Fatal(server.Run(":" + config.ServerPort))
}
