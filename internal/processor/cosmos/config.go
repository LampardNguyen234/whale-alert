package cosmos

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/processor/cosmos/staking"
	"github.com/LampardNguyen234/whale-alert/internal/processor/cosmos/transfer"
)

type CosmosProcessorConfig struct {
	Transfer transfer.TransferProcessorConfig `json:"Transfer"`
	Staking  staking.StakingProcessorConfig   `json:"Staking"`
}

func DefaultConfig() CosmosProcessorConfig {
	return CosmosProcessorConfig{
		Transfer: transfer.DefaultConfig(),
		Staking:  staking.DefaultConfig(),
	}
}

func (cfg CosmosProcessorConfig) IsValid() (bool, error) {
	if _, err := cfg.Transfer.IsValid(); err != nil {
		return false, fmt.Errorf("invalid transfer: %v", err)
	}
	if _, err := cfg.Staking.IsValid(); err != nil {
		return false, fmt.Errorf("invalid staking: %v", err)
	}

	return true, nil
}
