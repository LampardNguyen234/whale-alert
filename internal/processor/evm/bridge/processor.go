package bridge

import (
	"context"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
	"github.com/LampardNguyen234/whale-alert/internal/processor/evm/bridge/contracts"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
	"sync"
)

type BridgeProcessor struct {
	*processorCommon.BaseProcessor
	queue  chan interface{}
	cfg    BridgeConfig
	suites map[uint8]*BridgeClient
}

func NewBridgeProcessor(cfg BridgeConfig,
	db *store.Store,
	log logger.Logger,
) (*BridgeProcessor, error) {
	if _, err := cfg.IsValid(); err != nil {
		return nil, err
	}
	suites := make(map[uint8]*BridgeClient)
	for _, bCfg := range cfg.Chains {
		s, err := NewBridgeSuite(bCfg, log)
		if err != nil {
			return nil, err
		}
		suites[s.DomainID()] = s
	}

	return &BridgeProcessor{
		BaseProcessor: &processorCommon.BaseProcessor{
			Db:  db,
			Log: log.WithPrefix("Bridge Processor"),
			Mtx: new(sync.Mutex),
		},
		queue:  make(chan interface{}),
		cfg:    cfg,
		suites: suites,
	}, nil
}

func (p *BridgeProcessor) Queue(msg interface{}) {
	if _, ok := msg.(EventMsg); !ok {
		return
	}

	p.queue <- msg
}

func (p *BridgeProcessor) Start(ctx context.Context) {
	p.Log.Infof("STARTED")

	for _, client := range p.suites {
		go client.ListenToTxs(ctx, p.queue, nil)
	}
	for {
		select {
		case <-ctx.Done():
			p.Log.Infof("STOPPED")
			return
		case receipt := <-p.queue:
			//jsb, _ := json.Marshal(receipt)
			//p.Log.Debugf("message received: %v", string(jsb))
			msg, ok := receipt.(EventMsg)
			if !ok {
				continue
			}
			err := p.Process(ctx, msg)
			if err != nil {
				p.Log.Errorf("failed to process receipt %v: %v", receipt, err)
			}
		default:
		}
	}
}

func (p *BridgeProcessor) Process(ctx context.Context, msg EventMsg) error {
	switch msg.Event.Name() {
	case contracts.DepositEventName:
		return p.processDeposit(ctx, msg)
	case contracts.RelayerAddedEventName:
		return p.processRelayerAdded(ctx, msg)
	case contracts.RetryEventName:
		return p.processRetry(ctx, msg)
	case contracts.RoleGrantedEventName:
		return p.processRoleGranted(ctx, msg)
	case contracts.RoleRevokedEventName:
		return p.processRoleRevoked(ctx, msg)
	}

	return nil
}
