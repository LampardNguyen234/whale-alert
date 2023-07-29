package contracts

import (
	"fmt"
	evmCommon "github.com/LampardNguyen234/whale-alert/internal/processor/evm/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"strings"
)

// BaseContract holds the least information for a smart contract.
type BaseContract struct {
	ABI     abi.ABI
	Address common.Address
}

// BridgeContract holds the logic of a Bridge contract.
type BridgeContract struct {
	BaseContract
	Bridge
}

// NewBridgeContract creates a new BridgeContract given the necessary information.
func NewBridgeContract(address string, backEnd bind.ContractBackend) (*BridgeContract, error) {
	addr := common.HexToAddress(address)

	contractABI, err := abi.JSON(strings.NewReader(BridgeMetaData.ABI))
	if err != nil {
		return nil, err
	}

	bridge, err := NewBridge(addr, backEnd)
	if err != nil {
		return nil, err
	}

	return &BridgeContract{
		BaseContract: BaseContract{ABI: contractABI, Address: addr},
		Bridge:       *bridge,
	}, nil
}

func (c *BridgeContract) UnpackLog(log types.Log) (evmCommon.EVMEvent, error) {
	bridgeABI := c.ABI
	switch log.Topics[0] {
	case bridgeABI.Events[DepositEventName].ID:
		return c.Bridge.ParseDeposit(log)
	case bridgeABI.Events[RelayerThresholdChangedEventName].ID:
		return c.Bridge.ParseRelayerThresholdChanged(log)
	case bridgeABI.Events[RelayerAddedEventName].ID:
		return c.Bridge.ParseRelayerAdded(log)
	case bridgeABI.Events[RelayerRemovedEventName].ID:
		return c.Bridge.ParseRelayerRemoved(log)
	case bridgeABI.Events[ProposalEventName].ID:
		return c.Bridge.ParseProposalEvent(log)
	case bridgeABI.Events[ProposalVoteEventName].ID:
		return c.Bridge.ParseProposalVote(log)
	case bridgeABI.Events[FailedHandlerExecutionEventName].ID:
		return c.Bridge.ParseFailedHandlerExecution(log)
	case bridgeABI.Events[RetryEventName].ID:
		return c.Bridge.ParseRetry(log)
	case bridgeABI.Events[RoleGrantedEventName].ID:
		return c.Bridge.ParseRoleGranted(log)
	case bridgeABI.Events[RoleRevokedEventName].ID:
		return c.Bridge.ParseRoleRevoked(log)
	default:
		return nil, fmt.Errorf("log topic %v not supported", log.Topics[0])
	}
}

// Erc20HandlerContract holds the logic of an Erc20Handler contract.
type Erc20HandlerContract struct {
	BaseContract
	Erc20Handler
}

// NewERC20HandlerContract creates a new Erc20HandlerContract given the necessary information.
func NewERC20HandlerContract(address string, backEnd bind.ContractBackend) (*Erc20HandlerContract, error) {
	addr := common.HexToAddress(address)

	contractABI, err := abi.JSON(strings.NewReader(Erc20HandlerMetaData.ABI))
	if err != nil {
		return nil, err
	}

	erc20Handler, err := NewErc20Handler(addr, backEnd)
	if err != nil {
		return nil, err
	}

	return &Erc20HandlerContract{
		BaseContract: BaseContract{ABI: contractABI, Address: addr},
		Erc20Handler: *erc20Handler,
	}, nil
}

// GenericHandlerContract holds the logic of an Erc721Handler contract.
type GenericHandlerContract struct {
	BaseContract
	GenericHandler
}

// NewGenericHandlerContract creates a new GenericHandlerContract given the necessary information.
func NewGenericHandlerContract(address string, backEnd bind.ContractBackend) (*GenericHandlerContract, error) {
	addr := common.HexToAddress(address)

	contractABI, err := abi.JSON(strings.NewReader(GenericHandlerMetaData.ABI))
	if err != nil {
		return nil, err
	}

	genericHandler, err := NewGenericHandler(common.HexToAddress(address), backEnd)
	if err != nil {
		return nil, err
	}

	return &GenericHandlerContract{
		BaseContract:   BaseContract{ABI: contractABI, Address: addr},
		GenericHandler: *genericHandler,
	}, nil
}
