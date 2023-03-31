package clients

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/clients/cosmos"
	"github.com/LampardNguyen234/whale-alert/internal/clients/evm"
	"github.com/LampardNguyen234/whale-alert/internal/clients/tiki"
)

type ClientsConfig struct {
	Evm          evm.EvmClientConfig       `json:"Evm"`
	Cosmos       cosmos.CosmosClientConfig `json:"Cosmos"`
	TikiExchange tiki.TikiClientConfig     `json:"TikiExchange"`
}

func DefaultConfig() ClientsConfig {
	return ClientsConfig{
		Evm:          evm.DefaultConfig(),
		Cosmos:       cosmos.DefaultConfig(),
		TikiExchange: tiki.DefaultConfig(),
	}
}

// IsValid checks if the current ClientsConfig is valid.
func (cfg ClientsConfig) IsValid() (bool, error) {
	if _, err := cfg.Evm.IsValid(); err != nil {
		return false, fmt.Errorf("invalid Evm: %v", err)
	}

	if _, err := cfg.Cosmos.IsValid(); err != nil {
		return false, fmt.Errorf("invalid Cosmos: %v", err)
	}

	if _, err := cfg.TikiExchange.IsValid(); err != nil {
		return false, fmt.Errorf("invalid TikiExchange: %v", err)
	}

	return true, nil
}
