package tiki

import (
	"encoding/json"
	"testing"
)

func TestTikiClient_GetOrders(t *testing.T) {
	orders, err := c.GetOrders(ctx)
	if err != nil {
		panic(err)
	}

	jsb, _ := json.MarshalIndent(orders, "", "\t")
	c.log.Infof("%v", string(jsb))
}

func TestTikiClient_LatestOrderID(t *testing.T) {
	id, err := c.LatestOrderID(ctx)
	if err != nil {
		panic(err)
	}

	c.log.Infof("latestID: %v", id)
}
