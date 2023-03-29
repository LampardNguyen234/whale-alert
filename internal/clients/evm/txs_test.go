package evm

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestEvmClient_TransactionReceipt(t *testing.T) {
	receipt, err := c.TransactionReceipt(ctx, common.HexToHash("0xe030621498492ac16ccddb94b2cd634554621a5088617ed0498a9fe65f21081a"))
	if err != nil {
		panic(err)
	}

	jsb, _ := json.MarshalIndent(receipt, "", "\t")
	c.log.Infof(string(jsb))
}

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
