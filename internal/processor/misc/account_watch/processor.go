package account_watcher

import (
	"context"
	"github.com/LampardNguyen234/whale-alert/internal/clients/cosmos"
	"github.com/LampardNguyen234/whale-alert/internal/clients/evm"
	"github.com/LampardNguyen234/whale-alert/internal/common"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/core/types"
	evmTypes "github.com/evmos/ethermint/x/evm/types"
	lru "github.com/hashicorp/golang-lru"
	"sync"
	"time"
)

type AccountWatchProcessor struct {
	*processorCommon.BaseProcessor
	*evm.EvmClient
	*cosmos.CosmosClient

	queue          chan interface{}
	cfg            AccountWatchProcessorConfig
	cachedAccounts *lru.Cache
}

func NewAccountWatchProcessor(cfg AccountWatchProcessorConfig,
	evmClient *evm.EvmClient,
	cosmosClient *cosmos.CosmosClient,
	db *store.Store,
	log logger.Logger,
) (*AccountWatchProcessor, error) {
	cached, err := lru.New(1024)
	if err != nil {
		return nil, err
	}

	return &AccountWatchProcessor{
		BaseProcessor: &processorCommon.BaseProcessor{
			Db:  db,
			Log: log.WithPrefix("Account-Watcher"),
			Mtx: new(sync.Mutex),
		},
		EvmClient:      evmClient,
		CosmosClient:   cosmosClient,
		queue:          make(chan interface{}, cfg.QueueSize),
		cfg:            cfg,
		cachedAccounts: cached,
	}, nil
}

func (p *AccountWatchProcessor) Queue(msg interface{}) {
	evmReceipt, ok := msg.(*types.Receipt)
	if ok {
		p.Mtx.Lock()
		defer p.Mtx.Unlock()
		p.queue <- evmReceipt
		return
	}
	cosmosReceipt, ok := msg.(*sdk.TxResponse)
	if ok {
		p.Mtx.Lock()
		defer p.Mtx.Unlock()
		p.queue <- cosmosReceipt
		return
	}
}

func (p *AccountWatchProcessor) Start(ctx context.Context) {
	err := p.loadMonitoredAccounts()
	if err != nil {
		p.Log.Errorf("failed to cached monitored account: %v", err)
	}
	p.Log.Infof("STARTED")
	for {
		select {
		case <-ctx.Done():
			p.Log.Infof("STOPPED")
			return
		case receipt := <-p.queue:
			err = p.Process(ctx, receipt)
			if err != nil {
				p.Log.Errorf("failed to process receipt %v: %v", receipt, err)
			}
		default:
			time.Sleep(common.DefaultSleepTime)
		}
	}
}

func (p *AccountWatchProcessor) Process(ctx context.Context, receipt interface{}) error {
	if evmReceipt, ok := receipt.(*types.Receipt); ok {
		if p.EvmClient != nil {
			return p.processEVMTxs(ctx, evmReceipt)
		}
	} else if cosmosReceipt, ok := receipt.(*sdk.TxResponse); ok {
		if p.CosmosClient != nil {
			return p.processCosmosTxs(ctx, cosmosReceipt)
		}
	}

	return nil
}

func (p *AccountWatchProcessor) processEVMTxs(ctx context.Context, receipt *types.Receipt) error {
	if receipt.Status == types.ReceiptStatusSuccessful && !p.IsTxProcessed(receipt.TxHash.String()) {
		tx, _, err := p.EvmClient.TransactionByHash(ctx, receipt.TxHash)
		if err != nil {
			p.Log.Errorf("failed to get TransactionByHash %v: %v", receipt.TxHash, err)
			return err
		}

		to := tx.To().String()
		from := processorCommon.MustGetEvmTxSender(tx)
		from, to, direction, isMonitored := p.getTxMonitoredDetails(from, to)
		if !isMonitored {
			return nil
		}
		acc := to
		if direction == directionOut {
			acc = from
		}

		return p.Whm.Alert(Msg{
			TxMsg: processorCommon.TxMsg{
				From:   from,
				To:     to,
				Amount: common.FormatAmount(common.GetNormalizedValue(tx.Value())),
				TxHash: receipt.TxHash.String(),
			},
			Account:   acc,
			Type:      p.getEvmMsgType(receipt, tx),
			Direction: direction,
		}.String())
	}

	return nil
}

func (p *AccountWatchProcessor) processCosmosTxs(_ context.Context, receipt *sdk.TxResponse) error {
	if receipt.Code == 0 && !p.IsTxProcessed(receipt.TxHash) {
		messages := receipt.GetTx().GetMsgs()
		for _, msg := range messages {
			if _, ok := msg.(*evmTypes.MsgEthereumTx); ok {
				continue
			}

			from := p.CosmosClient.ParseCosmosMsgSender(msg)
			to := p.CosmosClient.ParseCosmosMsgReceiver(msg)
			from, to, direction, isMonitored := p.getTxMonitoredDetails(from, to)
			if !isMonitored {
				return nil
			}
			acc := to
			if direction == directionOut {
				acc = from
			}

			err := p.Whm.Alert(Msg{
				TxMsg: processorCommon.TxMsg{
					TxHash: receipt.TxHash,
					Amount: common.FormatAmount(p.CosmosClient.ParseCosmosMsgValue(msg)),
				},
				Type:      p.getCosmosMsgType(msg),
				Direction: direction,
				Account:   acc,
			}.String())
			if err != nil {
				return err
			}
		}
	}

	return nil
}
