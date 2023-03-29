package common

import (
	"context"
	"encoding/hex"
	"github.com/LampardNguyen234/whale-alert/common"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
	"github.com/LampardNguyen234/whale-alert/webhook"
	"strings"
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

	return accountDetail.String()
}

func (p *BaseProcessor) IsTxProcessed(txHashStr string) bool {
	txHashStr = strings.Replace(strings.ToLower(txHashStr), "0x", "", -1)
	txHash, err := hex.DecodeString(txHashStr)
	if err != nil {
		p.Log.Errorf("failed to decode transaction %v: %v", txHashStr, err)
		return false
	}

	processed, _ := p.Db.IsTxProcessed(txHash)
	return processed
}
