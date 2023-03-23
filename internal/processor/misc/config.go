package misc

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/processor/misc/account_watch"
)

type MiscProcessorConfig struct {
	AccountWatch account_watcher.AccountWatchProcessorConfig `json:"AccountWatch"`
}

func DefaultConfig() MiscProcessorConfig {
	return MiscProcessorConfig{
		AccountWatch: account_watcher.DefaultConfig(),
	}
}

func (cfg MiscProcessorConfig) IsValid() (bool, error) {
	if _, err := cfg.AccountWatch.IsValid(); err != nil {
		return false, fmt.Errorf("invalid accountWatch: %v", err)
	}

	return true, nil
}
