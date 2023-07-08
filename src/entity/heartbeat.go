package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Heartbeat struct {
	ID           uuid.UUID  `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key;"`
	ResourceID   uuid.UUID  `json:"resourceId" gorm:"type:uuid;"`
	StatusCode   int        `json:"statusCode" gorm:"type:integer;"`
	ResponseBody string     `json:"responseBody" gorm:"type:text;"`
	Latency      int64      `json:"latency" gorm:"type:bigint;"`
	HTTPTiming   HTTPTiming `json:"httpTiming" gorm:"type:jsonb"`
	CreatedAt    time.Time  `json:"createdAt"`
}

type HTTPTiming struct {
	DNSLookupTime     int64 `json:"dnsLookupTime"`
	ConnectionTime    int64 `json:"connectionTime"`
	TLSHandshakeTime  int64 `json:"tlsHandshakeTime"`
	FirstByteWaitTime int64 `json:"firstByteWaitTime"`
	TotalTime         int64 `json:"totalTime"`
}

// HTTPTiming Marshal
func (jsonField HTTPTiming) Value() (driver.Value, error) {
	return json.Marshal(jsonField)
}

// HTTPTiming Unmarshal
func (jsonField *HTTPTiming) Scan(value interface{}) error {
	data, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(data, &jsonField)
}

type Tabler interface {
	TableName() string
}

func (Heartbeat) TableName() string {
	return "resource_heartbeats"
}

func NewHeartbeat(res Resource, response *http.Response, latencyInMs int64, responseBody []byte, httpTiming HTTPTiming) Heartbeat {
	heartbeat := Heartbeat{}
	heartbeat.StatusCode = response.StatusCode
	heartbeat.ResponseBody = string(responseBody)
	heartbeat.ResourceID = res.ID
	heartbeat.Latency = latencyInMs
	heartbeat.HTTPTiming = httpTiming
	return heartbeat
}
