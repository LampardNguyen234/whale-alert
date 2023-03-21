package evm

import "fmt"

type EvmClientConfig struct {
	Endpoint string `json:"Endpoint"`
	Enabled  bool   `json:"Enabled"`
}

func DefaultConfig() EvmClientConfig {
	return EvmClientConfig{
		Endpoint: "http://127.0.0.1:8544",
		Enabled:  true,
	}
}

// IsValid checks if the current EvmClientConfig is valid.
func (cfg EvmClientConfig) IsValid() (bool, error) {
	if cfg.Endpoint == "" {
		return false, fmt.Errorf("empty endpoint")
	}

	return true, nil
}
