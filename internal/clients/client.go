package clients

import (
	"context"
	"github.com/LampardNguyen234/whale-alert/internal/clients/common"
	"github.com/LampardNguyen234/whale-alert/internal/clients/cosmos"
	"github.com/LampardNguyen234/whale-alert/internal/clients/evm"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
	"math/big"
)

type Client interface {
	ListenToTxs(ctx context.Context, resultChan chan interface{}, fromBlk *big.Int)
}

// NewClientsFromConfig creates new Client's from the given config.
func NewClientsFromConfig(cfg ClientsConfig, store *store.Store, log logger.Logger) (map[string]Client, error) {
	ret := make(map[string]Client)
	if cfg.Evm.Enabled {
		evmClient, err := evm.NewEvmClient(cfg.Evm, store, log)
		if err != nil {
			return nil, err
		}
		ret[common.EvmClientName] = evmClient
	}
	if cfg.Cosmos.Enabled {
		cosmosClient, err := cosmos.NewCosmosClient(cfg.Cosmos, store, log)
		if err != nil {
			return nil, err
		}
		ret[common.CosmosClientName] = cosmosClient
	}

	return ret, nil
}
