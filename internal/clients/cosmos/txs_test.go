package cosmos

import (
	"math/big"
	"testing"
)

func TestCosmosClient_ListenToTxs(t *testing.T) {
	receiptChan := make(chan interface{})
	go c.ListenToTxs(ctx, receiptChan, new(big.Int).SetUint64(1204500))

	for {
		select {
		case receipt := <-receiptChan:
			_ = receipt
			//c.log.Infof("%v", receipt)
		}
	}
}
