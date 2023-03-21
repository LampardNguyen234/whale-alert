package cosmos

import (
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
	"google.golang.org/grpc"
)
import "github.com/cosmos/cosmos-sdk/codec"

type CosmosClient struct {
	*grpc.ClientConn
	BankClient
	log   logger.Logger
	store *store.Store
}

// NewCosmosClient creates a new cosmos client.
func NewCosmosClient(cfg CosmosClientConfig, store *store.Store, log logger.Logger) (*CosmosClient, error) {
	grpcClient, err := grpc.Dial(
		cfg.Endpoint,
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	)
	if err != nil {
		return nil, err
	}

	return &CosmosClient{
		store:      store,
		log:        log.WithPrefix("Cosmos-client"),
		ClientConn: grpcClient,
		BankClient: NewBankClient(grpcClient),
	}, nil
}

func (c *CosmosClient) MakeCodec() codec.Codec {
	cdc := codec.NewAminoCodec(codec.NewLegacyAmino())

	return cdc
}
