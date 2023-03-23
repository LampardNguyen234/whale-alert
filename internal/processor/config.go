package processor

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/processor/cosmos"
	evmTransfer "github.com/LampardNguyen234/whale-alert/internal/processor/evm/transfer"
	"github.com/LampardNguyen234/whale-alert/internal/processor/misc"
)

type ProcessorsConfig struct {
	EvmTransfer evmTransfer.TransferProcessorConfig `json:"EvmTransfer"`
	Cosmos      cosmos.CosmosProcessorConfig        `json:"Cosmos"`
	Misc        misc.MiscProcessorConfig            `json:"Misc"`
}

func DefaultConfig() ProcessorsConfig {
	return ProcessorsConfig{
		EvmTransfer: evmTransfer.DefaultConfig(),
		Cosmos:      cosmos.DefaultConfig(),
		Misc:        misc.DefaultConfig(),
	}
}

func (cfg ProcessorsConfig) IsValid() (bool, error) {
	if _, err := cfg.EvmTransfer.IsValid(); err != nil {
		return false, fmt.Errorf("invalid EvmTransfer: %v", err)
	}
	if _, err := cfg.Cosmos.IsValid(); err != nil {
		return false, fmt.Errorf("invalid Cosmos: %v", err)
	}
	if _, err := cfg.Misc.IsValid(); err != nil {
		return false, fmt.Errorf("invalid Misc: %v", err)
	}

	return true, nil
}
