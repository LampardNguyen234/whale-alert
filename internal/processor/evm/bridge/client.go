package bridge

import (
	"context"
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/processor/evm/bridge/contracts"
	"github.com/LampardNguyen234/whale-alert/internal/processor/evm/common"
	"github.com/LampardNguyen234/whale-alert/logger"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	"time"
)

const (
	DefaultAdminRole   = "DefaultAdmin"
	BridgeRelayerRole  = "BridgeRelayer"
	BridgeRetrierRole  = "BridgeRetrier"
	BridgeOperatorRole = "BridgeOperator"
)

// BridgeClient holds all the contract of a Bridge.
type BridgeClient struct {
	Client   *ethclient.Client
	Bridge   *contracts.BridgeContract
	Cfg      contracts.BridgeNetworkConfig
	log      logger.Logger
	domainID uint8
	roles    map[ethCommon.Hash]string
}

// NewBridgeSuite creates a new BridgeClient with the given parameters.
func NewBridgeSuite(cfg contracts.BridgeNetworkConfig, log logger.Logger) (*BridgeClient, error) {
	backEnd, err := ethclient.Dial(cfg.URL)
	if err != nil {
		return nil, err
	}

	bridge, err := contracts.NewBridgeContract(cfg.Bridge, backEnd)
	if err != nil {
		return nil, err
	}
	domain, err := bridge.DomainID(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	return &BridgeClient{
		Client:   backEnd,
		Cfg:      cfg,
		Bridge:   bridge,
		log:      log.WithPrefix(fmt.Sprintf("Bridge %v Client", cfg.Name)),
		domainID: domain,
	}, nil
}

func (c *BridgeClient) DomainID() uint8 {
	return c.domainID
}

// ListenToTxs retrieves all events for the current bridge.
func (c *BridgeClient) ListenToTxs(ctx context.Context, resultChan chan interface{}, _ *big.Int) {
	c.log.Debugf("STARTED")
	startBlock := new(big.Int).SetUint64(c.Cfg.FromHeight)
	endBlock := big.NewInt(0)
	for {
		select {
		case <-ctx.Done():
			c.log.Debugf("STOPPED with err: %v", ctx.Err())
			return
		default:
			latestBlk, err := c.Client.BlockNumber(ctx)
			if err != nil {
				c.log.Errorf("failed to get latest block: %v", err)
				continue
			}

			if startBlock == nil || startBlock.Uint64() == 0 {
				startBlock = big.NewInt(int64(latestBlk))
			}
			if latestBlk <= startBlock.Uint64() {
				time.Sleep(3 * time.Second)
				continue
			}

			endBlock.Add(startBlock, new(big.Int).SetUint64(c.Cfg.BlockInterval))
			if endBlock.Uint64() > latestBlk {
				endBlock = new(big.Int).SetUint64(latestBlk)
			}
			if endBlock.Uint64() <= startBlock.Uint64() {
				continue
			}

			events, err := c.fetchEvents(ctx, ethereum.FilterQuery{
				FromBlock: startBlock,
				ToBlock:   new(big.Int).Sub(endBlock, big.NewInt(1)),
			})
			if err != nil {
				c.log.Errorf("failed to fetch event: %v", err)
				if !strings.Contains(err.Error(), "Too Many Requests") {
					c.log.Errorf("failed to fetch event: %v", err)
				} else {
					time.Sleep(3 * time.Second)
				}
				time.Sleep(1 * time.Second)
				continue
			}
			if len(events) > 0 {
				c.log.Debugf("found %v event(c) for blocks [%v, %v)\n",
					len(events),
					startBlock.Uint64(), endBlock.Uint64(),
				)
			}

			for _, event := range events {
				resultChan <- EventMsg{DomainID: c.DomainID(), Event: event}
			}
			c.log.Debugf("blk: %v, %v\n", startBlock.Uint64(), endBlock.Uint64())

			startBlock = new(big.Int).SetUint64(endBlock.Uint64())
		}
	}
}

func (c *BridgeClient) FormatTxHashLink(txHash string) string {
	if c.Cfg.Explorer == "" {
		return txHash
	}
	return fmt.Sprintf("%v/tx/%v", c.Cfg.Explorer, txHash)
}

func (c *BridgeClient) FormatAddressLink(address string) string {
	if c.Cfg.Explorer == "" {
		return address
	}
	return fmt.Sprintf("%v/address/%v", c.Cfg.Explorer, address)
}

func (c *BridgeClient) Roles() map[ethCommon.Hash]string {
	if c.roles == nil {
		tmp, err := c.bridgeRoles()
		if err != nil {
			c.log.Errorf("failed to get bridgeRoles: %v", err)
			return map[ethCommon.Hash]string{}
		}
		return tmp
	}

	return c.roles
}

func (c *BridgeClient) bridgeRoles() (map[ethCommon.Hash]string, error) {
	ret := make(map[ethCommon.Hash]string)
	role, err := c.Bridge.DEFAULTADMINROLE(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	ret[role] = DefaultAdminRole

	role, err = c.Bridge.RELAYERROLE(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	ret[role] = BridgeRelayerRole

	role, err = c.Bridge.RETRIERROLE(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	ret[role] = BridgeRelayerRole

	role, err = c.Bridge.RELAYERROLE(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	ret[role] = BridgeRetrierRole

	role, err = c.Bridge.OPERATORROLE(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	ret[role] = BridgeOperatorRole

	c.roles = ret
	return ret, nil
}

// fetchEvents retrieves all events matching the given query.
func (c *BridgeClient) fetchEvents(ctx context.Context, q ethereum.FilterQuery) ([]common.EVMEvent, error) {
	if q.Addresses == nil {
		q.Addresses = make([]ethCommon.Address, 0)
	}
	q.Addresses = append(q.Addresses, ethCommon.HexToAddress(c.Cfg.Bridge))

	logs, err := c.Client.FilterLogs(ctx, q)
	if err != nil {
		return nil, err
	}

	events := make([]common.EVMEvent, 0)
	for _, log := range logs {
		if log.Removed {
			continue
		}

		e, _ := c.Bridge.UnpackLog(log)
		if e == nil {
			continue
		}
		events = append(events, e)
	}
	return events, nil
}
