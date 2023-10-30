package cosmos

import (
	goClient "github.com/LampardNguyen234/astra-go-sdk/client"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
	"github.com/cosmos/cosmos-sdk/client"
	"google.golang.org/grpc"
)

type BaseClient struct {
	client.Context
	grpc *grpc.ClientConn
}

type CosmosClient struct {
	*goClient.CosmosClient
	log   logger.Logger
	store *store.Store
}

// NewCosmosClient creates a new cosmos client.
func NewCosmosClient(cfg CosmosClientConfig, store *store.Store, log logger.Logger) (*CosmosClient, error) {
	if _, err := cfg.IsValid(); err != nil {
		return nil, err
	}

	cli, err := goClient.NewCosmosClient(cfg.CosmosClientConfig)
	if err != nil {
		return nil, err
	}

	tmpLog := log.WithPrefix("Cosmos-client")

	return &CosmosClient{
		CosmosClient: cli,
		log:          tmpLog,
		store:        store,
	}, nil
}
