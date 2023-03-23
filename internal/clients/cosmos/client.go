package cosmos

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/evmos/ethermint/encoding"
	"github.com/evmos/evmos/v6/app"
	"github.com/tendermint/tendermint/rpc/client/http"
	"google.golang.org/grpc"
)

type BaseClient struct {
	client.Context
	grpc *grpc.ClientConn
}

type CosmosClient struct {
	*BaseClient
	BankClient
	StakingClient
	log   logger.Logger
	store *store.Store
}

// NewCosmosClient creates a new cosmos client.
func NewCosmosClient(cfg CosmosClientConfig, store *store.Store, log logger.Logger) (*CosmosClient, error) {
	if _, err := cfg.IsValid(); err != nil {
		return nil, err
	}

	initSdkConfig(cfg)
	encCfg := encoding.MakeConfig(app.ModuleBasics)
	rpcHttp, err := http.New(fmt.Sprintf("%v:%v", cfg.Endpoint, cfg.TendermintPort), "/websocket")
	if err != nil {
		return nil, err
	}
	clientCtx := client.Context{}.WithClient(rpcHttp).
		WithCodec(encCfg.Marshaler).
		WithInterfaceRegistry(encCfg.InterfaceRegistry).
		WithTxConfig(encCfg.TxConfig).
		WithLegacyAmino(encCfg.Amino).
		WithChainID(cfg.ChainID)

	baseClient := &BaseClient{
		Context: clientCtx,
	}
	tmpLog := log.WithPrefix("Cosmos-client")

	return &CosmosClient{
		BaseClient:    baseClient,
		store:         store,
		log:           tmpLog,
		BankClient:    NewBankClient(baseClient.Context),
		StakingClient: NewStakingClient(baseClient.Context),
	}, nil
}
