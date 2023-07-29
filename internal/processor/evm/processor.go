package evm

import (
	"github.com/LampardNguyen234/whale-alert/internal/clients/evm"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
	"github.com/LampardNguyen234/whale-alert/internal/processor/evm/bridge"
	"github.com/LampardNguyen234/whale-alert/internal/processor/evm/transfer"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
)

func NewProcessors(cfg EvmConfig,
	client *evm.EvmClient,
	db *store.Store,
	log logger.Logger,
) ([]processorCommon.Processor, error) {
	ret := make([]processorCommon.Processor, 0)
	if cfg.Transfer.Enabled {
		transferProcessor, err := transfer.NewTransferProcessor(cfg.Transfer, client, db, log)
		if err != nil {
			return nil, err
		}
		ret = append(ret, transferProcessor)
	}

	if cfg.Bridge.Enabled {
		bridgeProcessor, err := bridge.NewBridgeProcessor(cfg.Bridge, db, log)
		if err != nil {
			return nil, err
		}
		ret = append(ret, bridgeProcessor)
	}

	return ret, nil
}
