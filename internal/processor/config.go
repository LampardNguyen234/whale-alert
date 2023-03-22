package processor

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/processor/cosmos"
	evmTransfer "github.com/LampardNguyen234/whale-alert/internal/processor/evm/transfer"
)

type ProcessorsConfig struct {
	EvmTransfer evmTransfer.TransferProcessorConfig `json:"EvmTransfer"`
	Cosmos      cosmos.CosmosProcessorConfig        `json:"Cosmos"`
}

func DefaultConfig() ProcessorsConfig {
	return ProcessorsConfig{
		EvmTransfer: evmTransfer.DefaultConfig(),
		Cosmos:      cosmos.DefaultConfig(),
	}
}

func (cfg ProcessorsConfig) IsValid() (bool, error) {
	if _, err := cfg.EvmTransfer.IsValid(); err != nil {
		return false, fmt.Errorf("invalid EvmTransfer: %v", err)
	}
	if _, err := cfg.Cosmos.IsValid(); err != nil {
		return false, fmt.Errorf("invalid Cosmos: %v", err)
	}

	return true, nil
}
