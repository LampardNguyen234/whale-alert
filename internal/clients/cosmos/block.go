package cosmos

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/big"
)

// LatestBlockHeight returns the latest block height from the current chain.
func (c *CosmosClient) LatestBlockHeight(ctx context.Context) (*big.Int, error) {
	ret, err := c.Client.Block(ctx, nil)
	if err != nil {
		return nil, err
	}

	return new(big.Int).SetUint64(uint64(ret.Block.Height)), nil
}

// BlockTxsByHeight retrieves the receipts of all transaction in a block given its height.
func (c *CosmosClient) BlockTxsByHeight(ctx context.Context, blk *big.Int) ([]*sdk.TxResponse, error) {
	height := blk.Int64()
	block, err := c.Client.Block(ctx, &height)
	if err != nil {
		return nil, err
	}

	res := make([]*sdk.TxResponse, 0)
	for _, tx := range block.Block.Txs {
		receipt, err := c.TxByHash(fmt.Sprintf("%X", tx.Hash()))
		if err != nil {
			c.log.Errorf("failed to get transaction receipt of %v: %v", tx.Hash(), err)
			return nil, err
		}

		res = append(res, receipt)
	}

	return res, nil
}
