package misc

import (
	"github.com/LampardNguyen234/whale-alert/internal/clients/cosmos"
	"github.com/LampardNguyen234/whale-alert/internal/clients/evm"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
	account_watcher "github.com/LampardNguyen234/whale-alert/internal/processor/misc/account_watch"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
)

func NewProcessors(cfg MiscProcessorConfig,
	evmClient *evm.EvmClient,
	cosmosClient *cosmos.CosmosClient,
	db *store.Store,
	log logger.Logger,
) ([]processorCommon.Processor, error) {
	ret := make([]processorCommon.Processor, 0)
	if cfg.AccountWatch.Enabled {
		processor, err := account_watcher.NewAccountWatchProcessor(cfg.AccountWatch,
			evmClient,
			cosmosClient,
			db,
			log,
		)
		if err != nil {
			return nil, err
		}
		ret = append(ret, processor)
	}

	return ret, nil
}
