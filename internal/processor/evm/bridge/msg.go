package bridge

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/processor/evm/common"
)

type EventMsg struct {
	DomainID uint8
	Event    common.EVMEvent
}

func (p *BridgeProcessor) formatNetwork(id uint8) string {
	return fmt.Sprintf("%v (%v)", p.suites[id].Cfg.Name, id)
}
