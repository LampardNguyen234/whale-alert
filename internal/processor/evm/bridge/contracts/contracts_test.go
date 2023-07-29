package contracts

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

func TestBridgeContract_UnpackLog(t *testing.T) {
	c, err := ethclient.Dial("https://rpc.ankr.com/bsc")
	if err != nil {
		panic(err)
	}

	bridge, err := NewBridgeContract("0x5fC4435AcA131f1F541D2fc67DC3A6a20d10a99d", c)
	if err != nil {
		panic(err)
	}

	receipt, err := c.TransactionReceipt(context.Background(), common.HexToHash("0xebddecc2453b7b174a27226190b6ee5b5e3a0ec7d968c2b825e6f0f51f4a1ff4"))
	if err != nil {
		panic(err)
	}

	for _, log := range receipt.Logs {
		e, err := bridge.UnpackLog(*log)
		if err != nil {
			fmt.Printf("%v, %v\n", log.Topics[0], err)
			continue
		}

		jsb, _ := json.Marshal(e)
		fmt.Println(string(jsb))
	}

}
