package evm

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestEvmClient_ListenToTxs(t *testing.T) {
	receiptChan := make(chan interface{})
	go c.ListenToTxs(ctx, receiptChan, nil)

	for {
		select {
		case receipt := <-receiptChan:
			jsb, _ := json.MarshalIndent(receipt, "", "\t")
			c.log.Infof("New receipt: %v", string(jsb))
		}
	}
}

func TestEvmClient_subscribePendingTxs(t *testing.T) {
	ret := make(chan common.Hash)
	go c.subscribePendingTxs(ctx, ret)

	for {
		select {
		case txHash := <-ret:
			c.log.Infof("New tx: %v", txHash)
		}
	}
}
