package tiki

import (
	"context"
	"fmt"
	"github.com/LampardNguyen234/whale-alert/common"
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

func (ord Order) Hash() common.Hash {
	return common.Digest(
		ord.ID,
		ord.Price,
		ord.Amount,
		ord.Total,
		ord.Side,
		ord.CreatedAt.String(),
	)
}

type Orders []*Order

// Amount returns the total amounts of the given Orders.
func (ord Orders) Amount() float64 {
	tmp := []*Order(ord)

	resp := float64(0)
	for _, tmpOrd := range tmp {
		resp += tmpOrd.Amount
	}
	return resp
}

func (ord Orders) ToSlice() []*Order {
	return ord
}

func (ord Orders) Count() int {
	return len(ord)
}

// AmountByType returns the total amounts of the given Orders by OrderType.
func (ord Orders) AmountByType(side OrderType) float64 {
	resp := make([]*Order, 0)
	ords := ord.ToSlice()
	for _, tmpOrd := range ords {
		if tmpOrd.Side == side {
			tmp := *tmpOrd
			resp = append(resp, &tmp)
		}
	}

	return Orders(resp).Amount()
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
