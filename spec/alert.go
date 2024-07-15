package spec

import (
	"time"
)

type AlertLevel string

const (
	Info     AlertLevel = "INFO"
	Warning  AlertLevel = "WARNING"
	Critical AlertLevel = "CRITICAL"
)

type AlertComponent string

const (
	APIServer AlertComponent = "API_Server"

	CoboIndexder AlertComponent = "Cobo_Indexder"
	ZuluIndexer  AlertComponent = "Zulu_Indexer"

	BridgeService AlertComponent = "Bridge_Service"
)

type Alert struct {
	Level     AlertLevel             `json:"level"`
	Message   string                 `json:"message"`
	Timestamp time.Time              `json:"timestamp"`
	Component AlertComponent         `json:"component"`
	HostName  string                 `json:"host_name"`
	Details   map[string]interface{} `json:"details,omitempty"`
}
