package transfer

import (
	"context"
	"github.com/LampardNguyen234/whale-alert/internal/clients/evm"
	"github.com/LampardNguyen234/whale-alert/internal/common"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
	"github.com/LampardNguyen234/whale-alert/webhook"
	"github.com/dustin/go-humanize"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"sync"
	"time"
)

type TransferProcessor struct {
	*processorCommon.BaseProcessor
	*evm.EvmClient
	queue chan *types.Receipt
	cfg   TransferProcessorConfig
}

func NewTransferProcessor(cfg TransferProcessorConfig,
	evmClient *evm.EvmClient,
	db *store.Store,
	log logger.Logger,
) (*TransferProcessor, error) {
	return &TransferProcessor{
		BaseProcessor: &processorCommon.BaseProcessor{
			Db:  db,
			Log: log,
			Mtx: new(sync.Mutex),
		},
		EvmClient: evmClient,
		queue:     make(chan *types.Receipt, cfg.QueueSize),
		cfg:       cfg,
	}, nil
}

func (p *TransferProcessor) SetWebHookManager(whm webhook.WebHookManager) {
	p.Whm = whm
}

func (p *TransferProcessor) Queue(msg interface{}) {
	receipt, ok := msg.(*types.Receipt)
	if !ok {
		return
	}

	p.Mtx.Lock()
	defer p.Mtx.Unlock()
	p.queue <- receipt
}

func (p *TransferProcessor) Process(ctx context.Context, receipt *types.Receipt) error {
	if receipt.Status != types.ReceiptStatusSuccessful {
		return nil
	}
	if len(receipt.Logs) != 0 {
		return nil
	}

	tx, _, err := p.EvmClient.TransactionByHash(ctx, receipt.TxHash)
	if err != nil {
		p.Log.Errorf("failed to get TransactionByHash %v: %v", receipt.TxHash, err)
		return err
	}

	amt := new(big.Float).SetInt(tx.Value())
	amt = amt.Quo(amt, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), common.AsaDecimalsBigInt, nil)))
	amtFloat, _ := amt.Float64()

	if triggerredAmt, ok := p.cfg.WhaleDefinition[common.AsaAddress]; ok && amtFloat > triggerredAmt {
		from := ""
		p.Log.Debugf("chainID: %v", tx.ChainId())
		signer, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
		if err != nil {
			p.Log.Errorf("failed to get sender of tx %v: %v", tx.Hash().String(), err)
		} else {
			from = signer.String()
		}

		return p.Whm.Alert(Msg{
			From:      from,
			To:        tx.To().String(),
			Amount:    humanize.FtoaWithDigits(amtFloat, 5),
			Token:     "0x",
			TokenName: "ASA",
			TxHash:    receipt.TxHash.String(),
		}.String())
	}

	return nil
}

func (p *TransferProcessor) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			p.Log.Infof("TransferProcessor STOPPED")
			return
		case receipt := <-p.queue:
			err := p.Process(ctx, receipt)
			if err != nil {
				p.Log.Errorf("failed to process receipt %v: %v", receipt, err)
			}
		default:
			time.Sleep(common.DefaultSleepTime)
		}
	}
}