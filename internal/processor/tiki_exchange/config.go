package tiki_exchange

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/processor/tiki_exchange/order_watch"
	"github.com/LampardNguyen234/whale-alert/internal/processor/tiki_exchange/volume_watch"
)

type TikiProcessorConfig struct {
	OrderWatch  order_watch.OrderWatchConfig   `json:"OrderWatch"`
	VolumeWatch volume_watch.VolumeWatchConfig `json:"VolumeWatch"`
}

func DefaultConfig() TikiProcessorConfig {
	return TikiProcessorConfig{
		OrderWatch:  order_watch.DefaultConfig(),
		VolumeWatch: volume_watch.DefaultConfig(),
	}
}

func (cfg TikiProcessorConfig) IsValid() (bool, error) {
	if _, err := cfg.OrderWatch.IsValid(); err != nil {
		return false, fmt.Errorf("invalid orderWatch: %v", err)
	}

	if _, err := cfg.VolumeWatch.IsValid(); err != nil {
		return false, fmt.Errorf("invalid volumeWatch: %v", err)
	}

	return true, nil
}
