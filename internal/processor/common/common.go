package common

import (
	"context"
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
