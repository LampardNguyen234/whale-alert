package tiki

import (
	"encoding/json"
	"fmt"
	"github.com/LampardNguyen234/whale-alert/common"
	"github.com/tendermint/tendermint/libs/rand"
	"testing"
	"time"
)

func randomOrder() *Order {
	price := float64(1 + rand.Uint64()%300)
	amount := float64(1 + rand.Uint64()%10000)
	side := OrderSellType
	if common.RandomHash().Bytes()[0]%2 == 0 {
		side = OrderBuyType
	}
	createdAt := time.Now().Add(time.Duration(rand.Int64()))
	return &Order{
		ID:        rand.Uint64(),
		Price:     price,
		Amount:    amount,
		Total:     price * amount,
		Side:      side,
		CreatedAt: createdAt,
	}
}

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

func TestOrders_Amount(t *testing.T) {
	orders := Orders([]*Order{
		randomOrder(),
		randomOrder(),
		randomOrder(),
		randomOrder(),
		randomOrder(),
		randomOrder(),
	})

	for i, order := range orders {
		jsb, _ := json.Marshal(order)
		fmt.Println(i, string(jsb))
	}

	fmt.Println(orders.Amount(), orders.Count(), orders.AmountByType(OrderSellType), orders.AmountByType(OrderBuyType))
}
