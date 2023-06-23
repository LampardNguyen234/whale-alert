package processor

import (
	"github.com/LampardNguyen234/whale-alert/internal/clients"
	"github.com/LampardNguyen234/whale-alert/internal/clients/common"
	"github.com/LampardNguyen234/whale-alert/internal/clients/cosmos"
	"github.com/LampardNguyen234/whale-alert/internal/clients/evm"
	"github.com/LampardNguyen234/whale-alert/internal/clients/tiki"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
	cosmosProcessor "github.com/LampardNguyen234/whale-alert/internal/processor/cosmos"
	evmTransfer "github.com/LampardNguyen234/whale-alert/internal/processor/evm/transfer"
	"github.com/LampardNguyen234/whale-alert/internal/processor/misc"
	tiki_exchange "github.com/LampardNguyen234/whale-alert/internal/processor/tiki_exchange"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
)

func NewProcessors(cfg ProcessorsConfig,
	blkClients map[string]clients.Client,
	db *store.Store,
	log logger.Logger,
) ([]processorCommon.Processor, error) {
	ret := make([]processorCommon.Processor, 0)

	var evmClient *evm.EvmClient
	var cosmosClient *cosmos.CosmosClient
	var tikiClient *tiki.TikiClient

	if client, ok := blkClients[common.EvmClientName]; ok {
		evmClient = client.(*evm.EvmClient)
		if cfg.EvmTransfer.Enabled {
			p, err := evmTransfer.NewTransferProcessor(cfg.EvmTransfer, evmClient, db, log)
			if err != nil {
				return nil, err
			}

			ret = append(ret, p)
		}

	}

	if client, ok := blkClients[common.CosmosClientName]; ok {
		cosmosClient = client.(*cosmos.CosmosClient)
		cosmosProcessors, err := cosmosProcessor.NewProcessors(cfg.Cosmos, cosmosClient, db, log)
		if err != nil {
			return nil, err
		}

		ret = append(ret, cosmosProcessors...)
	}

	if client, ok := blkClients[common.TikiExchangeClientName]; ok {
		tikiClient = client.(*tiki.TikiClient)
		tikiProcessors, err := tiki_exchange.NewProcessors(cfg.TikiExchange, tikiClient, db, log)
		if err != nil {
			return nil, err
		}

		ret = append(ret, tikiProcessors...)
	}

	miscProcessors, err := misc.NewProcessors(cfg.Misc, evmClient, cosmosClient, db, log)
	if err != nil {
		return nil, err
	}
	ret = append(ret, miscProcessors...)

	return ret, nil
}
