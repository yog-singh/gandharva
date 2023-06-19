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
	ResouceID    uuid.UUID `json:"resourceId" gorm:"type:uuid;"`
	StatusCode   int       `json:"statusCode" gorm:"type:integer;"`
	ResponseBody string    `json:"responseBody" gorm:"type:text;"`
	Latency      int       `json:"latency" gorm:"type:integer;"`
	CreatedAt    time.Time `json:"createdAt"`
}

func NewHeartbeat(res Resource, response http.Response) Heartbeat {
	heartbeat := Heartbeat{}
	heartbeat.StatusCode = response.StatusCode

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	heartbeat.ResponseBody = string(body)
	heartbeat.ResouceID = res.ID
	return heartbeat
}
