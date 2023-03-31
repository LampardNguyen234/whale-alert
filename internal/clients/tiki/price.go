package tiki

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

type AsaTicker struct {
	Open               string `json:"open"`
	Last               string `json:"last"`
	Low                string `json:"low"`
	High               string `json:"high"`
	Amount             string `json:"amount"`
	Volume             string `json:"vol"`
	AvgPrice           string `json:"avg_price"`
	PriceChangePercent string `json:"price_change_percent"`
}

type AsaSummary struct {
	At     string    `json:"at"`
	Ticker AsaTicker `json:"ticker"`
}

// GetAsaPrice returns the price of an ASA on the Tiki Exchange.
func (c *TikiClient) GetAsaPrice(ctx context.Context) (float64, error) {
	summary, err := c.GetAsaSummary(ctx)
	if err != nil {
		return 0, err
	}

	price, err := strconv.ParseFloat(summary.Ticker.Last, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse price `%v`", summary.Ticker.Last)
	}

	return price, nil
}

// GetAsaSummary returns the price summary of Astra on the Tiki Exchange.
func (c *TikiClient) GetAsaSummary(ctx context.Context) (*AsaSummary, error) {
	path := c.parseUrl("public/markets/astra/summary")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	var ret AsaSummary
	err = c.parseResponse(resp, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
