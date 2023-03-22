package cosmos

import "github.com/LampardNguyen234/whale-alert/logger"

var (
	c *CosmosClient
)

func init() {
	var err error
	c, err = NewCosmosClient(DefaultConfig(), nil, logger.NewZeroLogger(""))
	if err != nil {
		panic(err)
	}
}
