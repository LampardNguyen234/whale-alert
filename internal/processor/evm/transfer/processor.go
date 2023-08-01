package transfer

import (
	"context"
	sdkCommon "github.com/LampardNguyen234/astra-go-sdk/common"
	"github.com/LampardNguyen234/whale-alert/internal/clients/evm"
	"github.com/LampardNguyen234/whale-alert/internal/common"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/core/types"
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
			Log: log.WithPrefix("Evm-transfer"),
			Mtx: new(sync.Mutex),
		},
		EvmClient: evmClient,
		queue:     make(chan *types.Receipt, cfg.QueueSize),
		cfg:       cfg,
	}, nil
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

	amtFloat := sdkCommon.ParseAmountToDec(sdk.NewCoin(sdkCommon.BaseDenom, sdk.NewIntFromBigInt(tx.Value()))).MustFloat64()
	tokenDetail := p.Db.GetTokenDetail(common.ZeroAddress)
	p.Log.Debugf("newEvmTransfer: %v, %v/%v", tx.Hash(), amtFloat, tokenDetail.WhaleDefinition)
	if amtFloat >= tokenDetail.WhaleDefinition {
		from := ""
		p.Log.Debugf("chainID: %v", tx.ChainId())
		signer, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
		if err != nil {
			p.Log.Errorf("failed to get sender of tx %v: %v", tx.Hash().String(), err)
		} else {
			from = signer.String()
		}

		return p.Whm.Alert(Msg{
			processorCommon.TxMsg{
				From:      p.ParseAccountDetail(from),
				To:        p.ParseAccountDetail(tx.To().String()),
				Amount:    common.FormatAmount(amtFloat),
				Token:     tokenDetail.TokenAddress,
				TokenName: tokenDetail.TokenName,
				TxHash:    receipt.TxHash.String(),
			},
		}.String())
	}

	return nil
}
