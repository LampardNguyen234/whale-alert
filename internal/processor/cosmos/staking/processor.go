package staking

import (
	"context"
	"github.com/LampardNguyen234/whale-alert/internal/clients/cosmos"
	"github.com/LampardNguyen234/whale-alert/internal/common"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
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
		switch msg.(type) {
		case *stakingTypes.MsgDelegate:
			tmpMsgDelegate := msg.(*stakingTypes.MsgDelegate)
			go p.processMsgDelegate(ctx, receipt, tmpMsgDelegate)
		case *stakingTypes.MsgUndelegate:
			tmpMsgUndelegate := msg.(*stakingTypes.MsgUndelegate)
			go p.processMsgUndelegate(ctx, receipt, tmpMsgUndelegate)
		case *stakingTypes.MsgBeginRedelegate:
			tmpMsgRedelegate := msg.(*stakingTypes.MsgBeginRedelegate)
			go p.processMsgRedelegate(ctx, receipt, tmpMsgRedelegate)
		case *stakingTypes.MsgCreateValidator:
			tmpMsgCreateValidator := msg.(*stakingTypes.MsgCreateValidator)
			go p.processMsgCreateValidator(ctx, receipt, tmpMsgCreateValidator)
		default:
			continue
		}
	}

	return nil
}

func (p *StakingProcessor) processMsgDelegate(ctx context.Context, receipt *sdk.TxResponse, msg *stakingTypes.MsgDelegate) {
	amtFloat := common.GetNormalizedValue(msg.Amount.Amount.BigInt())
	tokenDetail := p.Db.GetTokenDetail(common.ZeroAddress)
	p.Log.Debugf("newMsgDelegate: %v, %v/%v", *msg, amtFloat, tokenDetail.WhaleDefinition)
	if amtFloat >= tokenDetail.WhaleDefinition {
		err := p.Whm.Alert(DelegateMsg{
			TxMsg: processorCommon.TxMsg{
				From:      p.ParseAccountDetail(msg.DelegatorAddress),
				Amount:    common.FormatAmount(amtFloat),
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
	amtFloat := common.GetNormalizedValue(msg.Amount.Amount.BigInt())
	tokenDetail := p.Db.GetTokenDetail(common.ZeroAddress)
	p.Log.Debugf("newMsgUnDelegate: %v, %v/%v", *msg, amtFloat, tokenDetail.WhaleDefinition)
	if amtFloat >= p.Db.GetTokenDetail(common.ZeroAddress).WhaleDefinition {
		err := p.Whm.Alert(UndelegateMsg{
			TxMsg: processorCommon.TxMsg{
				From:      p.ParseAccountDetail(msg.DelegatorAddress),
				Amount:    common.FormatAmount(amtFloat),
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

func (p *StakingProcessor) processMsgRedelegate(ctx context.Context, receipt *sdk.TxResponse, msg *stakingTypes.MsgBeginRedelegate) {
	amtFloat := common.GetNormalizedValue(msg.Amount.Amount.BigInt())
	tokenDetail := p.Db.GetTokenDetail(common.ZeroAddress)
	p.Log.Debugf("newMsgRedelegate: %v, %v/%v", *msg, amtFloat, tokenDetail.WhaleDefinition)
	if amtFloat >= p.Db.GetTokenDetail(common.ZeroAddress).WhaleDefinition {
		err := p.Whm.Alert(BeginRedelegateMsg{
			TxMsg: processorCommon.TxMsg{
				From:      p.ParseAccountDetail(msg.DelegatorAddress),
				Amount:    common.FormatAmount(amtFloat),
				Token:     "0x",
				TokenName: "ASA",
				TxHash:    receipt.TxHash,
			},
			FromValidator: p.getValidatorName(ctx, msg.ValidatorSrcAddress),
			ToValidator:   p.getValidatorName(ctx, msg.ValidatorDstAddress),
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
		Address:    parseValidatorDetail(msg.ValidatorAddress, msg.Description),
		Name:       msg.Description.Moniker,
		Commission: msg.Commission.Rate.MustFloat64(),
	}.String())
	if err != nil {
		p.Log.Errorf("failed to processMsgDelegate txHash %v: %v", receipt.TxHash, err)
	}
}
