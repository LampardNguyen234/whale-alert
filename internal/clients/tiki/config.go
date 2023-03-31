package tiki

import "fmt"

type TikiClientConfig struct {
	Endpoint string `json:"Endpoint"`
	Enabled  bool   `json:"Enabled"`
}

func DefaultConfig() TikiClientConfig {
	return TikiClientConfig{
		Endpoint: "https://api.tiki.vn/sandseel/api/v2",
		Enabled:  true,
	}
}

// IsValid checks if the current EvmClientConfig is valid.
func (cfg TikiClientConfig) IsValid() (bool, error) {
	if cfg.Enabled && cfg.Endpoint == "" {
		return false, fmt.Errorf("empty endpoint")
	}

	return true, nil
}
