package common

import "fmt"

type BaseConfig struct {
	Enabled   bool `json:"Enabled"`
	QueueSize int  `json:"QueueSize"`
}

func (cfg BaseConfig) IsValid() (bool, error) {
	if cfg.Enabled && cfg.QueueSize == 0 {
		return false, fmt.Errorf("empty queueSize")
	}

	return true, nil
}

func DefaultBaseConfig() BaseConfig {
	return BaseConfig{
		true,
		1024,
	}
}
