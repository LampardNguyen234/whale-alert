package volume_watch

import (
	"context"
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/clients/tiki"
	"github.com/pkg/errors"
	"sort"
	"time"
)

var (
	errEmptyOrder   = fmt.Errorf("empty order")
	errNoOrderFound = fmt.Errorf("no order found")
)

func (p *VolumeWatchProcessor) processVolumeCheck(ctx context.Context, check *VolumeCheck) {
	p.Log.Debugf("processVolumeCheck(%v) STARTED", check.Period.String())
	for {
		select {
		case <-ctx.Done():
			p.Log.Debugf("processVolumeCheck(%v) DONE", check.Period.String())
			return
		default:
			time.Sleep(2 * time.Second)
			from := time.Now().Add(-check.Period)
			orders, err := p.getOrdersFromTime(from)
			if err != nil {
				if errors.Is(err, errEmptyOrder) || errors.Is(err, errNoOrderFound) {
					time.Sleep(10 * time.Second)
					continue
				}
				p.Log.Errorf("failed to get orders from %v: %v", from, err)
				continue
			}
			tmpOrders := tiki.Orders(orders)
			if tmpOrders.Amount() >= check.Volume {
				msg := Msg{
					Orders: tmpOrders,
					Period: check.Period,
				}.String()
				err = p.Whm.Alert(msg)
				if err != nil {
					p.Log.Errorf("failed to alert message %v: %v", msg, err)
					continue
				}

				if check.Period >= 1*time.Hour {
					time.Sleep(check.Period / 2)
				} else {
					time.Sleep(check.Period - check.Period*10/100)
				}
			}
		}
	}
}

func (p *VolumeWatchProcessor) enqueueOrder(order *tiki.Order) {
	p.Mtx.Lock()
	defer p.Mtx.Unlock()

	if _, exists := p.cache.Get(order.Hash().String()); exists {
		return
	}

	p.orders = append([]*tiki.Order{order}, p.orders...)
	sort.SliceStable(p.orders, func(i, j int) bool {
		return p.orders[i].CreatedAt.After(p.orders[j].CreatedAt)
	})

	p.cache.SetDefault(order.Hash().String(), true)
}

func (p *VolumeWatchProcessor) getOrdersFromTime(from time.Time) ([]*tiki.Order, error) {
	// assume that p.orders has been sorted with descending CreatedAt.
	p.Mtx.Lock()
	defer p.Mtx.Unlock()

	if len(p.orders) == 0 {
		return nil, errEmptyOrder
	}
	if len(p.orders) == 1 {
		if p.orders[0].CreatedAt.After(from) {
			return p.orders, nil
		}
		return nil, errNoOrderFound
	}

	// find stop index
	for i := len(p.orders) - 1; i >= 0; i-- {
		if p.orders[i].CreatedAt.After(from) {
			return p.orders[:i+1], nil
		}
	}

	return nil, errNoOrderFound
}

func (p *VolumeWatchProcessor) cleanOrders() {
	p.Mtx.Lock()
	defer p.Mtx.Unlock()
	if len(p.orders) == 0 {
		return
	}
	oldest := time.Now().Add(-p.cfg.VolumeChecks[0].Period)
	idx := len(p.orders)
	for i := len(p.orders) - 1; i >= 0; i-- {
		if p.orders[i].CreatedAt.After(oldest) {
			idx = i + 1
			break
		}
	}
	p.orders = p.orders[:idx]
	p.Log.Debugf("FINISH cleaning orders before %v", oldest.String())
}
