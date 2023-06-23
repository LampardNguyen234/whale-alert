package volume_watch

import (
	"context"
	"github.com/LampardNguyen234/whale-alert/common"
	"github.com/LampardNguyen234/whale-alert/internal/clients/tiki"
	internalCommon "github.com/LampardNguyen234/whale-alert/internal/common"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
	"sync"
	"time"
)

type VolumeWatchProcessor struct {
	*processorCommon.BaseProcessor
	*tiki.TikiClient

	orders []*tiki.Order
	cache  common.Cache
	cfg    VolumeWatchConfig
}

func NewVolumeWatchProcessor(cfg VolumeWatchConfig,
	tikiClient *tiki.TikiClient,
	db *store.Store,
	log logger.Logger,
) (*VolumeWatchProcessor, error) {
	return &VolumeWatchProcessor{
		BaseProcessor: &processorCommon.BaseProcessor{
			Db:  db,
			Log: log.WithPrefix("Tiki-Volume-Watcher"),
			Mtx: new(sync.Mutex),
		},
		TikiClient: tikiClient,
		orders:     make([]*tiki.Order, 0),
		cache:      common.NewSimpleCache(),
		cfg:        cfg,
	}, nil
}

func (p *VolumeWatchProcessor) Queue(msg interface{}) {
	order, ok := msg.(*tiki.Order)
	if !ok {
		return
	}
	p.enqueueOrder(order)
}

func (p *VolumeWatchProcessor) Start(ctx context.Context) {
	for _, vc := range p.cfg.VolumeChecks {
		go p.processVolumeCheck(ctx, vc)
	}

	p.Log.Infof("STARTED")
	cleanTime := time.Now()
	for {
		select {
		case <-ctx.Done():
			p.Log.Infof("STOPPED")
			return
		default:
			if time.Now().After(cleanTime) {
				p.cleanOrders()
				cleanTime = time.Now().Add(30 * time.Minute)
			}
			time.Sleep(internalCommon.DefaultSleepTime)
		}
	}
}
