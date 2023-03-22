package evm

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// LatestBlockHeight returns the latest block height from the current chain.
func (c *EvmClient) LatestBlockHeight(ctx context.Context) (*big.Int, error) {
	ret, err := c.Client.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}

	return new(big.Int).SetUint64(ret), nil
}

// BlockTxsByHeight retrieves the receipts of all transaction in a block given its height.
func (c *EvmClient) BlockTxsByHeight(ctx context.Context, blk *big.Int) ([]*types.Receipt, error) {
	block, err := c.Client.BlockByNumber(ctx, blk)
	if err != nil {
		return nil, err
	}

	res := make([]*types.Receipt, 0)
	for _, tx := range block.Transactions() {
		receipt, err := c.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			c.log.Errorf("failed to get transaction receipt of %v: %v", tx.Hash(), err)
			return nil, err
		}

		res = append(res, receipt)
	}

	return res, nil
}
