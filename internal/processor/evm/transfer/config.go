package transfer

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/common"
)

type TransferProcessorConfig struct {
	common.BaseConfig
}

func DefaultConfig() TransferProcessorConfig {
	return TransferProcessorConfig{
		BaseConfig: common.BaseConfig{Enabled: true, QueueSize: 1024},
	}
}

func (cfg TransferProcessorConfig) IsValid() (bool, error) {
	if cfg.QueueSize == 0 {
		return false, fmt.Errorf("invalid QueueSize")
	}

	return true, nil
}
