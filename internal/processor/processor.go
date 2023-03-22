package processor

import (
	"github.com/LampardNguyen234/whale-alert/internal/clients"
	"github.com/LampardNguyen234/whale-alert/internal/clients/common"
	"github.com/LampardNguyen234/whale-alert/internal/clients/cosmos"
	"github.com/LampardNguyen234/whale-alert/internal/clients/evm"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
	cosmosProcessor "github.com/LampardNguyen234/whale-alert/internal/processor/cosmos"
	evmTransfer "github.com/LampardNguyen234/whale-alert/internal/processor/evm/transfer"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
)

func NewProcessors(cfg ProcessorsConfig,
	blkClients map[string]clients.Client,
	db *store.Store,
	log logger.Logger,
) ([]processorCommon.Processor, error) {
	ret := make([]processorCommon.Processor, 0)

	if client, ok := blkClients[common.EvmClientName]; ok {
		tmpClient := client.(*evm.EvmClient)
		if cfg.EvmTransfer.Enabled {
			p, err := evmTransfer.NewTransferProcessor(cfg.EvmTransfer, tmpClient, db, log)
			if err != nil {
				return nil, err
			}

			ret = append(ret, p)
		}

	}

	if client, ok := blkClients[common.CosmosClientName]; ok {
		tmpClient := client.(*cosmos.CosmosClient)
		cosmosProcessors, err := cosmosProcessor.NewProcessors(cfg.Cosmos, tmpClient, db, log)
		if err != nil {
			return nil, err
		}

		ret = append(ret, cosmosProcessors...)
	}

	return ret, nil
}
