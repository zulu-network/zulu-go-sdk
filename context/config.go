package context

import (
	"github.com/zulu-network/zulu-go-sdk/log"
)

const (
	KeyConfFile = "ZULU_CONF_FILE"
)

// SystemConfig config of baetyl system
type SystemConfig struct {
	Logger log.Config `yaml:"logger,omitempty" json:"logger,omitempty"`
}
