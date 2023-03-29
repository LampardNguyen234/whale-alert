package account_watcher

import (
	"github.com/LampardNguyen234/whale-alert/internal/common"
)

type AccountWatchProcessorConfig struct {
	common.BaseConfig
}

func DefaultConfig() AccountWatchProcessorConfig {
	return AccountWatchProcessorConfig{
		BaseConfig: common.BaseConfig{
			Enabled:   true,
			QueueSize: 1024,
		},
	}
}

func (cfg AccountWatchProcessorConfig) IsValid() (bool, error) {
	if _, err := cfg.BaseConfig.IsValid(); err != nil {
		return false, err
	}

	return true, nil
}
