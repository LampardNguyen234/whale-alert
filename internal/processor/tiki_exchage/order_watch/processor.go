package order_watch

import (
	"context"
	"github.com/LampardNguyen234/whale-alert/internal/clients/tiki"
	"github.com/LampardNguyen234/whale-alert/internal/common"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
	"sync"
	"time"
)

type OrderWatchProcessor struct {
	*processorCommon.BaseProcessor
	*tiki.TikiClient

	queue chan *tiki.Order
	cfg   OrderWatchConfig
}

func NewOrderWatchProcessor(cfg OrderWatchConfig,
	tikiClient *tiki.TikiClient,
	db *store.Store,
	log logger.Logger,
) (*OrderWatchProcessor, error) {
	return &OrderWatchProcessor{
		BaseProcessor: &processorCommon.BaseProcessor{
			Db:  db,
			Log: log.WithPrefix("Tiki-Order-Watcher"),
			Mtx: new(sync.Mutex),
		},
		TikiClient: tikiClient,
		queue:      make(chan *tiki.Order, cfg.QueueSize),
		cfg:        cfg,
	}, nil
}

func (p *OrderWatchProcessor) Queue(msg interface{}) {
	order, ok := msg.(*tiki.Order)
	if !ok {
		return
	}
	p.Mtx.Lock()
	defer p.Mtx.Unlock()
	p.queue <- order
}

func (p *OrderWatchProcessor) Start(ctx context.Context) {
	p.Log.Infof("STARTED")
	for {
		select {
		case <-ctx.Done():
			p.Log.Infof("STOPPED")
			return
		case order := <-p.queue:
			err := p.Process(ctx, order)
			if err != nil {
				p.Log.Errorf("failed to process order %v: %v", order, err)
			}
		default:
			time.Sleep(common.DefaultSleepTime)
		}
	}
}

func (p *OrderWatchProcessor) Process(ctx context.Context, order *tiki.Order) error {
	if order.Amount < p.cfg.MinAmount {
		return nil
	}

	currentPrice, err := p.TikiClient.GetAsaPrice(ctx)
	if err != nil {
		p.Log.Errorf("failed to GetAsaPrice: %v", err)
		return err
	}
	return p.Whm.Alert(Msg{
		Order:        *order,
		CurrentPrice: currentPrice,
	}.String())
}
