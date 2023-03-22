package evm_transfer

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/common"
)

type TransferProcessorConfig struct {
	common.BaseConfig
	QueueSize       uint               `json:"QueueSize"`
	WhaleDefinition map[string]float64 `json:"WhaleDefinition"`
}

func DefaultConfig() TransferProcessorConfig {
	return TransferProcessorConfig{
		BaseConfig: common.BaseConfig{Enabled: true},
		QueueSize:  1024,
		WhaleDefinition: map[string]float64{
			common.AsaAddress: 100,
		},
	}
}

func (cfg TransferProcessorConfig) IsValid() (bool, error) {
	if cfg.QueueSize == 0 {
		return false, fmt.Errorf("invalid QueueSize")
	}

	return true, nil
}