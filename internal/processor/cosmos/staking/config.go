package staking

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/common"
)

type StakingProcessorConfig struct {
	common.BaseConfig
}

func DefaultConfig() StakingProcessorConfig {
	return StakingProcessorConfig{
		BaseConfig: common.BaseConfig{Enabled: true, QueueSize: 1024},
	}
}

func (cfg StakingProcessorConfig) IsValid() (bool, error) {
	if cfg.QueueSize == 0 {
		return false, fmt.Errorf("invalid QueueSize")
	}

	return true, nil
}
