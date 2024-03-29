package evm

import (
	"context"
	clientCommon "github.com/LampardNguyen234/whale-alert/internal/clients/common"
	internalCommon "github.com/LampardNguyen234/whale-alert/internal/common"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

var (
	blockRetryInterval = 3 * time.Second
)

func (c *EvmClient) ListenToTxs(ctx context.Context, txResult chan interface{}, startBlk *big.Int) {
	txHashChan := make(chan common.Hash)
	go c.subscribePendingTxs(ctx, txHashChan)

	var currentBlk *big.Int
	if startBlk != nil {
		currentBlk = new(big.Int).SetUint64(startBlk.Uint64())
	}
	for {
		select {
		case <-ctx.Done():
			c.log.Infof("ListenToTxs STOPPED")
			return
		case txHash := <-txHashChan:
			c.log.Infof("new pending transaction: %v", txHash)
			receipt, err := c.TransactionReceipt(ctx, txHash)
			if err != nil {
				c.log.Errorf("failed to get transactionReceipt of %v: %v", txHash, err)
			}
			txResult <- receipt
		default:
			head, err := c.LatestBlockHeight(ctx)
			if err != nil {
				c.log.Error("Unable to get latest block")
				time.Sleep(internalCommon.DefaultSleepTime)
				continue
			}
			if currentBlk == nil || currentBlk.Cmp(new(big.Int).SetUint64(0)) <= 0 {
				currentBlk = big.NewInt(head.Int64())
			}
			if head.Cmp(currentBlk) < 0 {
				time.Sleep(internalCommon.DefaultSleepTime)
				continue
			}

			txs, err := c.BlockTxsByHeight(ctx, currentBlk)
			if err != nil {
				c.log.Errorf("failed to get blockTxsByHeight(%v): %v", currentBlk.Uint64(), err)
				time.Sleep(internalCommon.DefaultSleepTime)
				continue
			}
			for _, tx := range txs {
				txResult <- tx
			}

			err = c.store.StoreLastBlock(currentBlk, clientCommon.EvmClientID)
			if err != nil {
				c.log.Errorf("failed to storeLastBlock(%v): %v", currentBlk.Uint64(), err)
			}

			if currentBlk.Uint64()%100 == 0 {
				c.log.Debugf("ListenToTxs finished block %v", currentBlk.Uint64())
			}
			currentBlk = currentBlk.Add(currentBlk, big.NewInt(1))
		}
	}
}

func (c *EvmClient) subscribePendingTxs(ctx context.Context, txResult chan common.Hash) {
	sub, err := c.GEthClient.SubscribePendingTransactions(ctx, txResult)
	if err != nil {
		c.log.Errorf("failed to subscribe to pending transaction: %v", err)
		return
	}
	c.log.Infof("Start listening to pending txs...")
	for {
		select {
		case <-ctx.Done():
			c.log.Infof("SubscribePendingTransactions STOPPED!")
			return
		case err = <-sub.Err():
			c.log.Errorf("Subscription error: %v", err)
			return
		}
	}
}
