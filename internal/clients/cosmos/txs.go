package cosmos

import (
	"context"
	"math/big"
)

type TxsClient struct {
}

func (c *CosmosClient) ListenToTxs(ctx context.Context, txResult chan interface{}, startBlk *big.Int) {
	return
}
