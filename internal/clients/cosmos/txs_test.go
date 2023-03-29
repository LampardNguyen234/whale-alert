package cosmos

import (
	"fmt"
	"math/big"
	"testing"
)

func TestTxsClient_TxByHash(t *testing.T) {
	resp, err := c.TxByHash("E8675D20EBC79CE0485FC8A1EC2BC90A5D68F60E5BF86D330D2A543DA854B4BF")
	if err != nil {
		panic(err)
	}

	c.log.Infof(fmt.Sprintf("%v", resp))
}

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
