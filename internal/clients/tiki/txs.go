package tiki

import (
	"context"
	"github.com/LampardNguyen234/whale-alert/internal/common"
	"math/big"
	"time"
)

func (c *TikiClient) newOrders(ctx context.Context, currentID uint64) ([]Order, error) {
	orders, err := c.GetOrders(ctx)
	if err != nil {
		return nil, err
	}

	ret := make([]Order, 0)
	for _, order := range orders {
		if order.ID > currentID {
			ret = append(ret, order)
		}
	}

	return ret, nil
}

func (c *TikiClient) ListenToTxs(ctx context.Context, txResult chan interface{}, _ *big.Int) {
	currentID := uint64(0)
	for {
		select {
		case <-ctx.Done():
			c.log.Infof("ListenToTxs STOPPED")
			return
		default:
			latestOrderID, err := c.LatestOrderID(ctx)
			if err != nil {
				c.log.Errorf("failed to get LatestOrderID: %v", err)
				time.Sleep(common.DefaultSleepTime)
				continue
			}
			if currentID == 0 {
				currentID = latestOrderID
			}
			if currentID == latestOrderID {
				time.Sleep(common.DefaultSleepTime)
				continue
			}

			newOrders, err := c.newOrders(ctx, currentID)
			if err != nil {
				c.log.Errorf("failed to get newOrders(%v): %v", currentID, err)
				time.Sleep(common.DefaultSleepTime)
				continue
			}
			for _, order := range newOrders {
				txResult <- &order
			}

			currentID = latestOrderID
		}
	}
}
