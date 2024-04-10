package context

import (
	"github.com/baetyl/baetyl-go/v2/log"
)

const (
	KeyConfFile = "ZULU_CONF_FILE"
)

// SystemConfig config of baetyl system
type SystemConfig struct {
	Logger log.Config `yaml:"logger,omitempty" json:"logger,omitempty"`
}
