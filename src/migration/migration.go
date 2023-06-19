package main

import (
	"fmt"
	"log"

	"github.com/yog-singh/gandharva/src/config"
	"github.com/yog-singh/gandharva/src/db"
	"github.com/yog-singh/gandharva/src/entity"
)

func init() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	db.ConnectDB(&config)
}

func main() {
	db.DB.AutoMigrate(&entity.Resource{})
	db.DB.AutoMigrate(&entity.Heartbeat{})
	fmt.Println("? Migration complete")
}
