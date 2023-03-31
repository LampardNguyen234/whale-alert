package tiki

import (
	"encoding/json"
	"testing"
)

func TestTikiClient_ListenToTxs(t *testing.T) {
	orderChan := make(chan interface{})
	go c.ListenToTxs(ctx, orderChan, nil)

	for {
		select {
		case order := <-orderChan:
			jsb, _ := json.Marshal(order)
			c.log.Infof("newOrder: %v", string(jsb))
		}
	}
}
