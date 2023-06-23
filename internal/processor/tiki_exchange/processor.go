package tiki_exchange

import (
	"github.com/LampardNguyen234/whale-alert/internal/clients/tiki"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
	"github.com/LampardNguyen234/whale-alert/internal/processor/tiki_exchange/order_watch"
	"github.com/LampardNguyen234/whale-alert/internal/processor/tiki_exchange/volume_watch"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
)

func NewProcessors(cfg TikiProcessorConfig,
	tikiClient *tiki.TikiClient,
	db *store.Store,
	log logger.Logger,
) ([]processorCommon.Processor, error) {
	ret := make([]processorCommon.Processor, 0)
	if cfg.OrderWatch.Enabled {
		processor, err := order_watch.NewOrderWatchProcessor(cfg.OrderWatch,
			tikiClient,
			db,
			log,
		)
		if err != nil {
			return nil, err
		}
		ret = append(ret, processor)
	}
	if cfg.VolumeWatch.Enabled {
		processor, err := volume_watch.NewVolumeWatchProcessor(cfg.VolumeWatch,
			tikiClient,
			db,
			log,
		)
		if err != nil {
			return nil, err
		}
		ret = append(ret, processor)
	}

	return ret, nil
}
