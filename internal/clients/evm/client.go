package evm

import (
	"context"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

type EvmClient struct {
	*ethclient.Client
	GEthClient *gethclient.Client
	RPCClient  *rpc.Client
	store      *store.Store
	log        logger.Logger
	chainID    *big.Int
}

// NewEvmClient creates a new EvmClient.
func NewEvmClient(cfg EvmClientConfig, store *store.Store, log logger.Logger) (*EvmClient, error) {
	rpcClient, err := rpc.Dial(cfg.Endpoint)
	if err != nil {
		return nil, err
	}

	client := ethclient.NewClient(rpcClient)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	return &EvmClient{
		Client:     client,
		GEthClient: gethclient.New(rpcClient),
		RPCClient:  rpcClient,
		store:      store,
		log:        log.WithPrefix("evm-client"),
		chainID:    chainID,
	}, nil
}

func (c *EvmClient) ChainID() *big.Int {
	return c.chainID
}
