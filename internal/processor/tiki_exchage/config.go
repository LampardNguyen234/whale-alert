package tiki_exchange

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/processor/tiki_exchage/order_watch"
)

type TikiProcessorConfig struct {
	OrderWatch order_watch.OrderWatchConfig `json:"OrderWatch"`
}

func DefaultConfig() TikiProcessorConfig {
	return TikiProcessorConfig{
		OrderWatch: order_watch.DefaultConfig(),
	}
}

func (cfg TikiProcessorConfig) IsValid() (bool, error) {
	if _, err := cfg.OrderWatch.IsValid(); err != nil {
		return false, fmt.Errorf("invalid orderWatch: %v", err)
	}

	return true, nil
}
