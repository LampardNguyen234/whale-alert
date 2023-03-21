package evm

import (
	"context"
	"github.com/LampardNguyen234/whale-alert/db"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
)

var (
	c   *EvmClient
	ctx context.Context
)

func init() {
	var err error
	tmpDb, err := db.NewLvlDB("./lvldbdata")
	if err != nil {
		panic(err)
	}
	s := store.NewStore(tmpDb)

	c, err = NewEvmClient(DefaultConfig(), s, logger.NewZeroLogger(""))
	if err != nil {
		panic(err)
	}
	ctx = context.Background()
}
