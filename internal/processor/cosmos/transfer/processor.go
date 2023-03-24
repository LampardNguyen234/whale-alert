package transfer

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/LampardNguyen234/whale-alert/internal/clients/cosmos"
	"github.com/LampardNguyen234/whale-alert/internal/common"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/dustin/go-humanize"
)

type TransferProcessor struct {
	*processorCommon.BaseProcessor
	*cosmos.CosmosClient
	queue chan *sdk.TxResponse
	cfg   TransferProcessorConfig
}

func NewTransferProcessor(cfg TransferProcessorConfig,
	cosmosClient *cosmos.CosmosClient,
	db *store.Store,
	log logger.Logger,
) (*TransferProcessor, error) {
	return &TransferProcessor{
		BaseProcessor: &processorCommon.BaseProcessor{
			Db:  db,
			Log: log.WithPrefix("Cosmos-transfer"),
			Mtx: new(sync.Mutex),
		},
		CosmosClient: cosmosClient,
		queue:        make(chan *sdk.TxResponse, cfg.QueueSize),
		cfg:          cfg,
	}, nil
}

func (p *TransferProcessor) Queue(msg interface{}) {
	receipt, ok := msg.(*sdk.TxResponse)
	if !ok {
		return
	}
	if receipt.Code != 0 {
		return
	}

	p.Mtx.Lock()
	defer p.Mtx.Unlock()
	p.queue <- receipt
}

func (p *TransferProcessor) Start(ctx context.Context) {
	p.Log.Infof("STARTED")
	for {
		select {
		case <-ctx.Done():
			p.Log.Infof("STOPPED")
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

func (p *TransferProcessor) Process(_ context.Context, receipt *sdk.TxResponse) error {
	messages := receipt.GetTx().GetMsgs()
	for _, msg := range messages {
		tmpMsg, ok := msg.(*bankTypes.MsgSend)
		if !ok {
			continue
		}

		amt := new(big.Float).SetInt(tmpMsg.Amount.AmountOf(cosmos.Denom).BigInt())
		amt = amt.Quo(amt, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), common.AsaDecimalsBigInt, nil)))
		amtFloat, _ := amt.Float64()

		if amtFloat >= p.cfg.MinAmount {
			return p.Whm.Alert(Msg{
				From:      p.ParseAccountDetail(tmpMsg.FromAddress),
				To:        p.ParseAccountDetail(tmpMsg.ToAddress),
				Amount:    humanize.FormatFloat("#,###.##",amtFloat),
				Token:     "0x",
				TokenName: "ASA",
				TxHash:    receipt.TxHash,
			}.String())
		}
	}

	return nil
}
