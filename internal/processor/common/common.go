package common

import (
	"context"
	"fmt"
	"github.com/LampardNguyen234/whale-alert/common"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
	"github.com/LampardNguyen234/whale-alert/webhook"
	"sync"
)

type Processor interface {
	Queue(interface{})
	Start(context.Context)
	SetWebHookManager(manager webhook.WebHookManager)
}

type BaseProcessor struct {
	Db  *store.Store
	Whm webhook.WebHookManager
	Log logger.Logger
	Mtx *sync.Mutex
}

func (p *BaseProcessor) SetWebHookManager(manager webhook.WebHookManager) {
	p.Whm = manager
}

func (p *BaseProcessor) ParseAccountDetail(addr string) string {
	addr = common.MustAccountAddressToHex(addr)

	accountDetail, err := p.Db.GetAccountDetail(addr)
	if err != nil {
		return addr
	}

	return fmt.Sprintf("%v (%v...%v)", accountDetail.Name, addr[:10], addr[len(addr)-10:])
}
