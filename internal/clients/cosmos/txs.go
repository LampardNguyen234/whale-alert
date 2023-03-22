package cosmos

import (
	"context"
	clientCommon "github.com/LampardNguyen234/whale-alert/internal/clients/common"
	"github.com/LampardNguyen234/whale-alert/internal/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"math/big"
	"time"
)

func (c *CosmosClient) TxByHash(hash string) (*sdk.TxResponse, error) {
	resp, err := tx.QueryTx(c.BaseClient.Context, hash)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CosmosClient) ListenToTxs(ctx context.Context, txResult chan interface{}, startBlk *big.Int) {
	for {
		select {
		case <-ctx.Done():
			c.log.Infof("ListenToTxs STOPPED")
			return
		default:
			head, err := c.LatestBlockHeight(ctx)
			if err != nil {
				c.log.Error("Unable to get latest block")
				time.Sleep(common.DefaultSleepTime)
				continue
			}
			if startBlk == nil || startBlk.Cmp(new(big.Int).SetUint64(0)) <= 0 {
				startBlk = big.NewInt(head.Int64())
			}
			if head.Cmp(startBlk) < 0 {
				time.Sleep(common.DefaultSleepTime)
				continue
			}

			txs, err := c.BlockTxsByHeight(ctx, startBlk)
			if err != nil {
				c.log.Errorf("failed to get blockTxsByHeight(%v): %v", startBlk.Uint64(), err)
				continue
			}
			for _, tmpTx := range txs {
				txResult <- tmpTx
			}

			err = c.store.StoreLastBlock(startBlk, clientCommon.CosmosClientID)
			if err != nil {
				c.log.Errorf("failed to storeLastBlock(%v): %v", startBlk.Uint64(), err)
			}

			if startBlk.Uint64()%10 == 0 {
				c.log.Debugf("ListenToTxs finished block %v", startBlk.Uint64())
			}
			startBlk = startBlk.Add(startBlk, big.NewInt(1))
		}
	}
}
