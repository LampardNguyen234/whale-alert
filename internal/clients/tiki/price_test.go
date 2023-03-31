package tiki

import (
	"encoding/json"
	"testing"
)

func TestTikiClient_GetAsaPrice(t *testing.T) {
	price, err := c.GetAsaPrice(ctx)
	if err != nil {
		panic(err)
	}

	c.log.Infof("price: %v", price)
}

func TestTikiClient_GetAsaSummary(t *testing.T) {
	summary, err := c.GetAsaSummary(ctx)
	if err != nil {
		panic(err)
	}

	jsb, _ := json.MarshalIndent(summary, "", "\t")

	c.log.Infof("%v", string(jsb))
}
