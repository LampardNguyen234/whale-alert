package order_watch

import (
	"github.com/LampardNguyen234/whale-alert/internal/common"
)

type OrderWatchConfig struct {
	common.BaseConfig
	MinAmount float64
}

func DefaultConfig() OrderWatchConfig {
	return OrderWatchConfig{
		BaseConfig: common.BaseConfig{
			Enabled:   true,
			QueueSize: 1024,
		},
		MinAmount: 1000,
	}
}

func (cfg OrderWatchConfig) IsValid() (bool, error) {
	if _, err := cfg.BaseConfig.IsValid(); err != nil {
		return false, err
	}

	return true, nil
}
