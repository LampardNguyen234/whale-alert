package cosmos

import "fmt"

type CosmosClientConfig struct {
	Endpoint       string `json:"Endpoint"`
	GrpcPort       string `json:"GrpcPort"`
	TendermintPort string `json:"TendermintPort"`
	Prefix         string `json:"Prefix"`
	ChainID        string `json:"ChainID"`
	Enabled        bool   `json:"Enabled"`
}

func DefaultConfig() CosmosClientConfig {
	return CosmosClientConfig{
		GrpcPort:       "9090",
		TendermintPort: "26657",
		Endpoint:       "http://206.189.158.191",
		Prefix:         "astra",
		ChainID:        "astra-11115_1",
		Enabled:        false,
	}
}

// IsValid checks if the current CosmosClientConfig is valid.
func (cfg CosmosClientConfig) IsValid() (bool, error) {
	if cfg.Endpoint == "" {
		return false, fmt.Errorf("empty endpoint")
	}
	if cfg.ChainID == "" {
		return false, fmt.Errorf("empty chainID")
	}
	if cfg.GrpcPort == "" {
		return false, fmt.Errorf("empty GrpcPort")
	}
	if cfg.TendermintPort == "" {
		return false, fmt.Errorf("empty TendermintPort")
	}

	return true, nil
}
