package listener

import (
	"context"
	"encoding/json"
	"github.com/LampardNguyen234/whale-alert/internal/clients"
	clientsCommon "github.com/LampardNguyen234/whale-alert/internal/clients/common"
	"github.com/LampardNguyen234/whale-alert/internal/common"
	"github.com/LampardNguyen234/whale-alert/internal/processor"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/LampardNguyen234/whale-alert/logger"
	"github.com/LampardNguyen234/whale-alert/webhook"
	"math/big"
	"sync"
	"time"
)

type Listener struct {
	clients    map[string]clients.Client
	processors []processor.Processor
	db         *store.Store
	log        logger.Logger
	whm        webhook.WebHookManager
	mtx        *sync.Mutex
	cfg        ListenerConfig
}

// NewListener creates a new listener.
func NewListener(cfg ListenerConfig, clients map[string]clients.Client, processors []processor.Processor, db *store.Store, log logger.Logger, whm webhook.WebHookManager) (*Listener, error) {
	return &Listener{
		clients:    clients,
		processors: processors,
		db:         db,
		log:        log.WithPrefix("Listener"),
		whm:        whm,
		mtx:        new(sync.Mutex),
		cfg:        cfg,
	}, nil
}

func (l *Listener) Start(ctx context.Context) {
	resultChan := make(chan interface{})

	var startBlock *big.Int
	switch l.cfg.StartBlock {
	case -1:
		latestBlk, err := l.db.GetLastBlock(clientsCommon.EvmClientID)
		if err != nil {
			l.log.Errorf("failed to getLastBlock: %v", err)
		}
		startBlock = latestBlk
	case 0:
		startBlock = nil
	default:
		startBlock = big.NewInt(l.cfg.StartBlock)
	}
	for _, p := range l.processors {
		p.SetWebHookManager(l.whm)
		go p.Start(ctx)
	}
	for _, c := range l.clients {
		go c.ListenToTxs(ctx, resultChan, startBlock)
	}
	for {
		select {
		case <-ctx.Done():
			l.log.Infof("STOPPED")
			return
		case msg := <-resultChan:
			if err, ok := msg.(error); ok {
				l.log.Errorf("new error received: %v", err)
				continue
			}
			for _, p := range l.processors {
				go p.Queue(msg)
			}
			jsb, _ := json.Marshal(msg)
			l.log.Infof("new msg: %v", string(jsb))
		default:
			time.Sleep(common.DefaultSleepTime)
		}
	}
}
