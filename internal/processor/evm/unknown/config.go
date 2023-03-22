package transfer

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/common"
)

type UnknownProcessorConfig struct {
	common.BaseConfig
	QueueSize       uint               `json:"QueueSize"`
	WhaleDefinition map[string]float64 `json:"WhaleDefinition"`
}

func DefaultConfig() UnknownProcessorConfig {
	return UnknownProcessorConfig{
		BaseConfig: common.BaseConfig{Enabled: true},
		QueueSize:  1024,
		WhaleDefinition: map[string]float64{
			common.AsaAddress: 100,
		},
	}
}

func (cfg UnknownProcessorConfig) IsValid() (bool, error) {
	if cfg.QueueSize == 0 {
		return false, fmt.Errorf("invalid QueueSize")
	}

	return true, nil
}
