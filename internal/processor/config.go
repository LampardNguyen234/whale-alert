package processor

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/processor/cosmos"
	"github.com/LampardNguyen234/whale-alert/internal/processor/evm"
	"github.com/LampardNguyen234/whale-alert/internal/processor/misc"
	"github.com/LampardNguyen234/whale-alert/internal/processor/tiki_exchange"
)

type ProcessorsConfig struct {
	Evm          evm.EvmConfig                     `json:"Evm"`
	Cosmos       cosmos.CosmosProcessorConfig      `json:"Cosmos"`
	Misc         misc.MiscProcessorConfig          `json:"Misc"`
	TikiExchange tiki_exchange.TikiProcessorConfig `json:"TikiExchange"`
}

func DefaultConfig() ProcessorsConfig {
	return ProcessorsConfig{
		Evm:          evm.DefaultConfig(),
		Cosmos:       cosmos.DefaultConfig(),
		Misc:         misc.DefaultConfig(),
		TikiExchange: tiki_exchange.DefaultConfig(),
	}
}

func (cfg ProcessorsConfig) IsValid() (bool, error) {
	if _, err := cfg.Evm.IsValid(); err != nil {
		return false, fmt.Errorf("invalid Evm: %v", err)
	}
	if _, err := cfg.Cosmos.IsValid(); err != nil {
		return false, fmt.Errorf("invalid Cosmos: %v", err)
	}
	if _, err := cfg.Misc.IsValid(); err != nil {
		return false, fmt.Errorf("invalid Misc: %v", err)
	}
	if _, err := cfg.TikiExchange.IsValid(); err != nil {
		return false, fmt.Errorf("invalid TikiExchange: %v", err)
	}

	return true, nil
}
