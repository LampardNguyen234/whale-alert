package cosmos

import "fmt"

type CosmosClientConfig struct {
	Endpoint string `json:"Endpoint"`
	Enabled  bool   `json:"Enabled"`
}

func DefaultConfig() CosmosClientConfig {
	return CosmosClientConfig{Endpoint: "http://127.0.0.1:26657", Enabled: false}
}

// IsValid checks if the current CosmosClientConfig is valid.
func (cfg CosmosClientConfig) IsValid() (bool, error) {
	if cfg.Endpoint == "" {
		return false, fmt.Errorf("empty endpoint")
	}

	return true, nil
}
