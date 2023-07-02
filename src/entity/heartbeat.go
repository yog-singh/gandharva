package entity

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Heartbeat struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key;"`
	ResourceID   uuid.UUID `json:"resourceId" gorm:"type:uuid;"`
	StatusCode   int       `json:"statusCode" gorm:"type:integer;"`
	ResponseBody string    `json:"responseBody" gorm:"type:text;"`
	Latency      int64     `json:"latency" gorm:"type:bigint;"`
	CreatedAt    time.Time `json:"createdAt"`
}

type Tabler interface {
	TableName() string
}

func (Heartbeat) TableName() string {
	return "resource_heartbeats"
}

func NewHeartbeat(res Resource, response *http.Response, latencyInMs int64) Heartbeat {
	heartbeat := Heartbeat{}
	heartbeat.StatusCode = response.StatusCode

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	heartbeat.ResponseBody = string(body)
	heartbeat.ResourceID = res.ID
	heartbeat.Latency = latencyInMs
	return heartbeat
}
