package volume_watch

import (
	"github.com/LampardNguyen234/whale-alert/internal/clients/tiki"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
)

type VolumeWatchProcessor struct {
	*processorCommon.BaseProcessor
	*tiki.TikiClient

	queue chan *tiki.Order
	cfg   VolumeWatchConfig
}
