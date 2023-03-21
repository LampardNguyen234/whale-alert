package processor

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/processor/evm_transfer"
)

type ProcessorsConfig struct {
	EvmTransfer evm_transfer.TransferProcessorConfig `json:"EvmTransfer"`
}

func DefaultConfig() ProcessorsConfig {
	return ProcessorsConfig{
		EvmTransfer: evm_transfer.DefaultConfig(),
	}
}

func (cfg ProcessorsConfig) IsValid() (bool, error) {
	if _, err := cfg.EvmTransfer.IsValid(); err != nil {
		return false, fmt.Errorf("invalid EvmTransfer: %v", err)
	}
	return true, nil
}
