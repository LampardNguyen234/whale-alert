package tiki

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"time"
)

type OrderType string

const (
	OrderSellType OrderType = "sell"
	OrderBuyType  OrderType = "buy"
)

const (
	DefaultOrderLimit = 100
)

type Order struct {
	ID        uint64    `json:"id"`
	Price     float64   `json:"price"`
	Amount    float64   `json:"amount"`
	Total     float64   `json:"total"`
	Side      OrderType `json:"side"`
	CreatedAt time.Time `json:"createAt"`
}

// GetOrders returns the current newest orders on the Tiki Exchange.
//
// If `limits` is not set, the default number of orders will be returned is DefaultOrderLimit.
func (c *TikiClient) GetOrders(ctx context.Context, limits ...int) ([]Order, error) {
	limit := DefaultOrderLimit
	if len(limits) > 0 {
		limit = limits[0]
	}
	path := c.parseUrl(fmt.Sprintf("public/markets/asaxu/trades?limit=%v", limit))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	type retHolder struct {
		*Order
		At        int64     `json:"created_at"`
		TakerType OrderType `json:"taker_type"`
	}

	tmpRet := make([]retHolder, 0)
	err = c.parseResponse(resp, &tmpRet)
	if err != nil {
		return nil, err
	}

	ret := make([]Order, 0)
	for _, order := range tmpRet {
		order.CreatedAt = time.Unix(order.At, 0)
		order.Side = OrderBuyType
		if order.TakerType == OrderSellType {
			order.Side = OrderSellType
		}
		ret = append(ret, *order.Order)
	}
	sort.SliceStable(ret, func(i, j int) bool {
		return ret[i].ID >= ret[j].ID
	})

	return ret, nil
}

func (c *TikiClient) LatestOrderID(ctx context.Context) (uint64, error) {
	orders, err := c.GetOrders(ctx, 1)
	if err != nil {
		return 0, err
	}
	if len(orders) == 0 {
		return 0, fmt.Errorf("no order found")
	}

	return orders[0].ID, nil
}
