package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/yog-singh/gandharva/src/db"
	"github.com/yog-singh/gandharva/src/entity"
	"github.com/yog-singh/gandharva/src/middleware"
	"gorm.io/gorm"
)

func AddResource(resource entity.Resource) (entity.Resource, error) {
	entity.InitializeResource(&resource)
	if result := db.DB.Create(&resource); result.Error != nil {
		fmt.Println(result.Error)
		return entity.Resource{}, result.Error
	}
	return resource, nil
}

func GetAllResources() ([]entity.Resource, error) {
	var resources []entity.Resource
	selectClause := `id, resource_id, status_code, latency, response_body, created_at`
	subQuery := `(select id, resource_id, status_code, latency, response_body, created_at, row_number() over(partition by resource_id order by created_at desc) as row_num from resource_heartbeats) resource_heartbeats`
	if result := db.DB.Preload("Heartbeats", func(db *gorm.DB) *gorm.DB {
		return db.Select(selectClause).Table(subQuery).Where("resource_heartbeats.id = id AND resource_heartbeats.row_num = 1")
	}).Find(&resources); result.Error != nil {
		fmt.Println(result.Error)
		return []entity.Resource{}, result.Error
	}
	return resources, nil
}

func CheckResourceHeartbeat() error {
	var resources []entity.Resource

	if result := db.DB.Where("((EXTRACT(EPOCH FROM (CURRENT_TIMESTAMP - last_checked_at)) / 60) > ping_interval_in_mins)").Find(&resources); result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	for _, resource := range resources {
		go getResponseAndSaveHeartbeat(resource)
	}

	return nil
}

func getResponseAndSaveHeartbeat(resource entity.Resource) {
	response := middleware.CheckHeartbeat(resource)
	heartbeat := entity.NewHeartbeat(resource, response.Response, response.RequestCompletionTimeInMs)
	db.DB.Create(&heartbeat)
	updateResourceStatus(&resource, response.Response)
	db.DB.Save(resource)
}

func updateResourceStatus(resource *entity.Resource, response *http.Response) {
	if response.StatusCode >= 200 && response.StatusCode <= 299 {
		resource.Status = entity.RESOURCE_UP
	} else if response.StatusCode >= 400 && response.StatusCode <= 499 {
		resource.Status = entity.RESOURCE_INACCESSIBLE
	} else if response.StatusCode >= 500 && response.StatusCode <= 599 {
		resource.Status = entity.RESOURCE_DOWN
	}
	resource.LastCheckedAt = time.Now()
}
