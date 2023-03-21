package common

import (
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
	"github.com/LampardNguyen234/whale-alert/webhook"
	"sync"
)

type BaseProcessor struct {
	Db  *store.Store
	Whm webhook.WebHookManager
	Log logger.Logger
	Mtx *sync.Mutex
}
