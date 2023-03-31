package cosmos

import "fmt"

type CosmosClientConfig struct {
	Endpoint       string `json:"Endpoint"`
	TendermintPort string `json:"TendermintPort"`
	Prefix         string `json:"Prefix"`
	ChainID        string `json:"ChainID"`
	Enabled        bool   `json:"Enabled"`
}

func DefaultConfig() CosmosClientConfig {
	return CosmosClientConfig{
		TendermintPort: "26657",
		Endpoint:       "http://127.0.0.1",
		Prefix:         "astra",
		ChainID:        "astra-11115_1",
		Enabled:        false,
	}
}

// IsValid checks if the current CosmosClientConfig is valid.
func (cfg CosmosClientConfig) IsValid() (bool, error) {
	if !cfg.Enabled {
		return true, nil
	}
	if cfg.Endpoint == "" {
		return false, fmt.Errorf("empty endpoint")
	}
	if cfg.ChainID == "" {
		return false, fmt.Errorf("empty chainID")
	}
	if cfg.TendermintPort == "" {
		return false, fmt.Errorf("empty TendermintPort")
	}

	return true, nil
}
