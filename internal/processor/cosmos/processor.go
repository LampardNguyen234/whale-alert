package cosmos

import (
	"github.com/LampardNguyen234/whale-alert/internal/clients/cosmos"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
	"github.com/LampardNguyen234/whale-alert/internal/processor/cosmos/staking"
	"github.com/LampardNguyen234/whale-alert/internal/processor/cosmos/transfer"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
)

func NewProcessors(cfg CosmosProcessorConfig,
	client *cosmos.CosmosClient,
	db *store.Store,
	log logger.Logger,
) ([]processorCommon.Processor, error) {
	ret := make([]processorCommon.Processor, 0)
	if cfg.Staking.Enabled {
		delegateProcessor, err := staking.NewStakingProcessor(cfg.Staking, client, db, log)
		if err != nil {
			return nil, err
		}
		ret = append(ret, delegateProcessor)
	}

	if cfg.Transfer.Enabled {
		transferProcessor, err := transfer.NewTransferProcessor(cfg.Transfer, client, db, log)
		if err != nil {
			return nil, err
		}
		ret = append(ret, transferProcessor)
	}

	return ret, nil
}
