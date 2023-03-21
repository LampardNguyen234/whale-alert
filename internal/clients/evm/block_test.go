package evm

import "testing"

func TestEvmClient_LatestBlock(t *testing.T) {
	latestBlk, err := c.LatestBlockHeight(ctx)
	if err != nil {
		panic(err)
	}

	c.log.Infof("LatestBlk: %v", latestBlk.Uint64())
}
