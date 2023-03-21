package processor

import (
	"context"
	"github.com/LampardNguyen234/whale-alert/internal/clients"
	"github.com/LampardNguyen234/whale-alert/internal/clients/common"
	"github.com/LampardNguyen234/whale-alert/internal/clients/evm"
	"github.com/LampardNguyen234/whale-alert/internal/processor/evm_transfer"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
	"github.com/LampardNguyen234/whale-alert/webhook"
)

type Processor interface {
	Queue(interface{})
	Start(context.Context)
	SetWebHookManager(manager webhook.WebHookManager)
}

func NewProcessors(cfg ProcessorsConfig,
	blkClients map[string]clients.Client,
	db *store.Store,
	log logger.Logger,
) ([]Processor, error) {
	ret := make([]Processor, 0)

	if client, ok := blkClients[common.EvmClientName]; ok {
		evmClient := client.(*evm.EvmClient)
		if cfg.EvmTransfer.Enabled {
			p, err := evm_transfer.NewTransferProcessor(cfg.EvmTransfer, evmClient, db, log)
			if err != nil {
				return nil, err
			}

			ret = append(ret, p)
		}

	}

	return ret, nil
}
