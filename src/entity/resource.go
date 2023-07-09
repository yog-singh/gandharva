package entity

import (
	"time"

	"github.com/google/uuid"
)

type Resource struct {
	ID                 uuid.UUID   `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key;"`
	Name               string      `json:"name" gorm:"type:varchar(255);not null;"`
	Url                string      `json:"url" gorm:"type:varchar(512);not null;"`
	RequestMethod      string      `json:"requestMethod" gorm:"type:varchar(32);"`
	PingIntervalInMins int         `json:"pingIntervalInMins" gorm:"type:integer;"`
	Status             string      `json:"status" gorm:"type:varchar(64);`
	ExpectedStatusCode int         `json:"expectedStatusCode" gorm:"type:integer;"`
	LastCheckedAt      time.Time   `json:"lastCheckedAt" gorm:"type:timestamp;"`
	CreatedAt          time.Time   `json:"createdAt" gorm:"type:timestamp;"`
	Heartbeats         []Heartbeat `json:"heartbeats" gorm:"foreignKey:ResourceID;"`
}

// Resource States
const (
	PENDING               = "PENDING"
	RESOURCE_UP           = "RESOURCE_UP"
	RESOURCE_DOWN         = "RESOURCE_DOWN"
	RESOURCE_INACCESSIBLE = "RESOURCE_INACCESSIBLE"
)

func InitializeResource(resource *Resource) {
	resource.Status = PENDING
}
