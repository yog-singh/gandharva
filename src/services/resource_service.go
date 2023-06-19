package services

import (
	"fmt"
	"time"

	"github.com/yog-singh/gandharva/src/db"
	"github.com/yog-singh/gandharva/src/entity"
	"github.com/yog-singh/gandharva/src/middleware"
)

func AddResource(resource entity.Resource) (entity.Resource, error) {
	if result := db.DB.Create(&resource); result.Error != nil {
		fmt.Println(result.Error)
		return entity.Resource{}, result.Error
	}
	return resource, nil
}

func GetAllResources() ([]entity.Resource, error) {
	var resources []entity.Resource

	if result := db.DB.Preload("Heartbeats", "(EXTRACT(epoch FROM (CURRENT_TIMESTAMP - created_at)) / 3600) < ?", 24).Find(&resources); result.Error != nil {
		fmt.Println(result.Error)
		return []entity.Resource{}, result.Error
	}
	return resources, nil
}

func PingResources() error {
	var resources []entity.Resource

	if result := db.DB.Find(&resources); result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	for _, resource := range resources {
		response := middleware.CheckHeartbeat(resource)
		heartbeat := entity.NewHeartbeat(resource, *response)
		db.DB.Create(&heartbeat)
		resource.LastCheckedAt = time.Now()
		db.DB.Save(resource)
	}

	return nil
}
