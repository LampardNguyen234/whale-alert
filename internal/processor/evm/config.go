package evm

import (
	"github.com/LampardNguyen234/whale-alert/internal/processor/evm/bridge"
	evmTransfer "github.com/LampardNguyen234/whale-alert/internal/processor/evm/transfer"
	"github.com/pkg/errors"
)

type EvmConfig struct {
	Transfer evmTransfer.TransferProcessorConfig `json:"Transfer"`
	Bridge   bridge.BridgeConfig                 `json:"Bridge"`
}

func DefaultConfig() EvmConfig {
	return EvmConfig{
		Transfer: evmTransfer.DefaultConfig(),
		Bridge:   bridge.DefaultConfig(),
	}
}

func (cfg *EvmConfig) IsValid() (bool, error) {
	if _, err := cfg.Transfer.IsValid(); err != nil {
		return false, errors.Wrapf(err, "invalid transfer")
	}

	if _, err := cfg.Bridge.IsValid(); err != nil {
		return false, errors.Wrapf(err, "invalid bridge")
	}

	return true, nil
}
