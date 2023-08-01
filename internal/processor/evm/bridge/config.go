package bridge

import (
	"github.com/LampardNguyen234/whale-alert/internal/common"
	"github.com/LampardNguyen234/whale-alert/internal/processor/evm/bridge/contracts"
	"github.com/pkg/errors"
)

type BridgeConfig struct {
	common.BaseConfig
	Chains map[string]contracts.BridgeNetworkConfig
}

func (cfg *BridgeConfig) IsValid() (bool, error) {
	if !cfg.Enabled {
		return true, nil
	}
	if _, err := cfg.BaseConfig.IsValid(); err != nil {
		return false, err
	}

	for network, tmp := range cfg.Chains {
		if _, err := tmp.IsValid(); err != nil {
			return false, errors.Wrapf(err, network)
		}
	}

	return true, nil
}

func DefaultConfig() BridgeConfig {
	return BridgeConfig{
		BaseConfig: common.DefaultBaseConfig(),
		Chains: map[string]contracts.BridgeNetworkConfig{
			"ASA": {
				Name:          "Astra",
				FromHeight:    0,
				BlockInterval: 0,
				Bridge:        "0xf188be7da55bd7b649b1f3a4eb4f038e0e87095f",
				Explorer:      "https://explorer.astranaut.io",
			},
			"BSC": {
				Name:          "Binance Smart Chain",
				FromHeight:    0,
				BlockInterval: 0,
				Bridge:        "0x5fC4435AcA131f1F541D2fc67DC3A6a20d10a99d",
				Explorer:      "https://bscscan.com",
			},
		},
	}
}
