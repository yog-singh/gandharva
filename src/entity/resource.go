package entity

import (
	"time"

	"github.com/google/uuid"
)

type Resource struct {
	ID            uuid.UUID   `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name          string      `json:"name" gorm:"type:varchar(255);not null"`
	Url           string      `json:"url" gorm:"type:varchar(512);not null"`
	Type          string      `json:"type" gorm:"type:varchar(32);"`
	LatencyInMins int         `json:"latencyInMins" gorm:"type:integer"`
	LastCheckedAt time.Time   `json:"lastCheckedAt" gorm:"type:timestamp"`
	CreatedAt     time.Time   `json:"createdAt" gorm:"type:timestamp;"`
	Heartbeats    []Heartbeat `json:"heartbeats" gorm:"foreignKey:ResouceID"`
}
