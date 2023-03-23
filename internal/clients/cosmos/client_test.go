package cosmos

import (
	"context"
	"github.com/LampardNguyen234/whale-alert/db"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
)

var (
	c   *CosmosClient
	ctx = context.Background()
)

func init() {
	tmpDb, err := db.NewLvlDB("./lvldbdata")
	if err != nil {
		panic(err)
	}
	s := store.NewStore(tmpDb)

	c, err = NewCosmosClient(DefaultConfig(), s, logger.NewZeroLogger(""))
	if err != nil {
		panic(err)
	}
}
