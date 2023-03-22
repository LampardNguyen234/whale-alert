package staking

import (
	"context"
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/clients/cosmos"
	"github.com/LampardNguyen234/whale-alert/internal/common"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/dustin/go-humanize"
	"math/big"
	"sync"
	"time"
)

type StakingProcessor struct {
	*processorCommon.BaseProcessor
	*cosmos.CosmosClient
	queue chan *sdk.TxResponse
	cfg   StakingProcessorConfig
}

func NewStakingProcessor(cfg StakingProcessorConfig,
	cosmosClient *cosmos.CosmosClient,
	db *store.Store,
	log logger.Logger,
) (*StakingProcessor, error) {
	return &StakingProcessor{
		BaseProcessor: &processorCommon.BaseProcessor{
			Db:  db,
			Log: log.WithPrefix("Cosmos-staking"),
			Mtx: new(sync.Mutex),
		},
		CosmosClient: cosmosClient,
		queue:        make(chan *sdk.TxResponse, cfg.QueueSize),
		cfg:          cfg,
	}, nil
}

func (p *StakingProcessor) Queue(msg interface{}) {
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

func (p *StakingProcessor) Start(ctx context.Context) {
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

func (p *StakingProcessor) Process(ctx context.Context, receipt *sdk.TxResponse) error {
	messages := receipt.GetTx().GetMsgs()
	for _, msg := range messages {
		tmpMsgDelegate, ok := msg.(*stakingTypes.MsgDelegate)
		if ok {
			go p.processMsgDelegate(ctx, receipt, tmpMsgDelegate)
			continue
		}

		tmpMsgUndelegate, ok := msg.(*stakingTypes.MsgUndelegate)
		if ok {
			go p.processMsgUndelegate(ctx, receipt, tmpMsgUndelegate)
		}

		tmpMsgCreateValidator, ok := msg.(*stakingTypes.MsgCreateValidator)
		if ok {
			go p.processMsgCreateValidator(ctx, receipt, tmpMsgCreateValidator)
		}
	}

	return nil
}

func (p *StakingProcessor) processMsgDelegate(ctx context.Context, receipt *sdk.TxResponse, msg *stakingTypes.MsgDelegate) {
	amt := new(big.Float).SetInt(msg.Amount.Amount.BigInt())
	amt = amt.Quo(amt, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), common.AsaDecimalsBigInt, nil)))
	amtFloat, _ := amt.Float64()

	if amtFloat >= p.cfg.MinAmount {
		err := p.Whm.Alert(DelegateMsg{
			TxMsg: processorCommon.TxMsg{
				From:      msg.DelegatorAddress,
				Amount:    humanize.FtoaWithDigits(amtFloat, 5),
				Token:     "0x",
				TokenName: "ASA",
				TxHash:    receipt.TxHash,
			},
			Validator: p.getValidatorName(ctx, msg.ValidatorAddress),
		}.String())
		if err != nil {
			p.Log.Errorf("failed to processMsgDelegate txHash %v: %v", receipt.TxHash, err)
		}
	}
}

func (p *StakingProcessor) processMsgUndelegate(ctx context.Context, receipt *sdk.TxResponse, msg *stakingTypes.MsgUndelegate) {
	amt := new(big.Float).SetInt(msg.Amount.Amount.BigInt())
	amt = amt.Quo(amt, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), common.AsaDecimalsBigInt, nil)))
	amtFloat, _ := amt.Float64()

	if amtFloat >= p.cfg.MinAmount {
		err := p.Whm.Alert(UndelegateMsg{
			TxMsg: processorCommon.TxMsg{
				From:      msg.DelegatorAddress,
				Amount:    humanize.FtoaWithDigits(amtFloat, 5),
				Token:     "0x",
				TokenName: "ASA",
				TxHash:    receipt.TxHash,
			},
			Validator: p.getValidatorName(ctx, msg.ValidatorAddress),
		}.String())
		if err != nil {
			p.Log.Errorf("failed to processMsgDelegate txHash %v: %v", receipt.TxHash, err)
		}
	}
}

func (p *StakingProcessor) processMsgCreateValidator(_ context.Context, receipt *sdk.TxResponse, msg *stakingTypes.MsgCreateValidator) {
	err := p.Whm.Alert(CreateValidatorMsg{
		TxMsg: processorCommon.TxMsg{
			TxHash: receipt.TxHash,
		},
		Address:    msg.ValidatorAddress,
		Name:       msg.Description.Moniker,
		Commission: msg.Commission.Rate.MustFloat64(),
	}.String())
	if err != nil {
		p.Log.Errorf("failed to processMsgDelegate txHash %v: %v", receipt.TxHash, err)
	}
}

func (p *StakingProcessor) getValidatorName(ctx context.Context, valAddr string) string {
	validatorDetail, err := p.CosmosClient.GetValidatorDetail(ctx, valAddr)
	if err == nil && validatorDetail.Description.Moniker != "" {
		return fmt.Sprintf("%v (%v...%v)", validatorDetail.Description.Moniker,
			valAddr[:12],
			valAddr[len(valAddr)-12:],
		)
	}

	return valAddr
}

func parseValidatorDetail(valAddr string, desc stakingTypes.Description) string {
	if desc.Moniker != "" {
		return fmt.Sprintf("%v (%v...%v)", desc.Moniker,
			valAddr[:12],
			valAddr[len(valAddr)-12:],
		)
	}

	return valAddr
}