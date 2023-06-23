package volume_watch

import (
	"encoding/json"
	"github.com/LampardNguyen234/whale-alert/common"
	"github.com/LampardNguyen234/whale-alert/internal/clients/tiki"
	"github.com/LampardNguyen234/whale-alert/logger"
	"github.com/tendermint/tendermint/libs/rand"
	"testing"
	"time"
)

var p *VolumeWatchProcessor

func randomOrder() *tiki.Order {
	price := float64(1 + rand.Uint64()%300)
	amount := float64(1 + rand.Uint64()%10000)
	side := tiki.OrderSellType
	if common.RandomHash().Bytes()[0]%2 == 0 {
		side = tiki.OrderBuyType
	}
	createdAt := time.Now().Add(-time.Duration(int64(rand.Uint64()%100)) * time.Second)
	return &tiki.Order{
		ID:        rand.Uint64(),
		Price:     price,
		Amount:    amount,
		Total:     price * amount,
		Side:      side,
		CreatedAt: createdAt,
	}
}

func init() {
	var err error

	p, err = NewVolumeWatchProcessor(DefaultConfig(), nil, nil, logger.NewZeroLogger(""))
	if err != nil {
		panic(err)
	}
}

func TestVolumeWatchProcessor_getOrdersFromTime(t *testing.T) {
	for i := 0; i < 10; i++ {
		ord := randomOrder()
		p.enqueueOrder(ord)
	}

	for i := 0; i < len(p.orders); i++ {
		jsb, _ := json.Marshal(p.orders[i])
		p.Log.Debugf("order %v: %v", i, string(jsb))
	}

	ords, err := p.getOrdersFromTime(time.Now().Add(-50 * time.Second))
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(ords); i++ {
		jsb, _ := json.Marshal(ords[i])
		p.Log.Debugf("newOrder %v: %v", i, string(jsb))
	}
}
