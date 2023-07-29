package contracts

import (
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	DepositEventName = "Deposit"

	RelayerThresholdChangedEventName = "RelayerThresholdChanged"

	RelayerAddedEventName = "RelayerAdded"

	RelayerRemovedEventName = "RelayerRemoved"

	ProposalEventName = "ProposalEvent"

	ProposalVoteEventName = "ProposalVote"

	FailedHandlerExecutionEventName = "FailedHandlerExecution"

	RetryEventName = "Retry"
)

const (
	RoleGrantedEventName = "RoleGranted"
	RoleRevokedEventName = "RoleRevoked"
)

const (
	// ProposalInactive represents the `Inactive` status of a proposal.
	ProposalInactive = uint8(iota)

	// ProposalActive represents the `Active` status of a proposal.
	ProposalActive

	// ProposalPassed represents the `Passed` status of a proposal.
	ProposalPassed

	// ProposalExecuted represents the `Executed` status of a proposal.
	ProposalExecuted

	// ProposalCancelled represents the `Cancelled` status of a proposal.
	ProposalCancelled
)

// Proposal wraps the `Proposal` struct of the bridge contract.
type Proposal struct {
	_Status        uint8
	_YesVotesTotal uint8
	_ProposedBlock uint
}

func (e *BridgeDeposit) Name() string {
	return DepositEventName
}

func (e *BridgeDeposit) GetLog() types.Log {
	return e.Raw
}

func (e *BridgeRelayerThresholdChanged) Name() string {
	return RelayerThresholdChangedEventName
}

func (e *BridgeRelayerThresholdChanged) GetLog() types.Log {
	return e.Raw
}

func (e *BridgeRelayerAdded) Name() string {
	return RelayerAddedEventName
}

func (e *BridgeRelayerAdded) GetLog() types.Log {
	return e.Raw
}

func (e *BridgeRelayerRemoved) Name() string {
	return RelayerRemovedEventName
}

func (e *BridgeRelayerRemoved) GetLog() types.Log {
	return e.Raw
}

func (e *BridgeProposalEvent) Name() string {
	return ProposalEventName
}

func (e *BridgeProposalEvent) GetLog() types.Log {
	return e.Raw
}

func (e *BridgeProposalVote) Name() string {
	return ProposalVoteEventName
}

func (e *BridgeProposalVote) GetLog() types.Log {
	return e.Raw
}

func (e *BridgeFailedHandlerExecution) Name() string {
	return FailedHandlerExecutionEventName
}

func (e *BridgeFailedHandlerExecution) GetLog() types.Log {
	return e.Raw
}

func (e *BridgeRetry) Name() string {
	return RetryEventName
}

func (e *BridgeRetry) GetLog() types.Log {
	return e.Raw
}

func (e *BridgeRoleGranted) Name() string {
	return RoleGrantedEventName
}

func (e *BridgeRoleGranted) GetLog() types.Log {
	return e.Raw
}

func (e *BridgeRoleRevoked) Name() string {
	return RoleRevokedEventName
}

func (e *BridgeRoleRevoked) GetLog() types.Log {
	return e.Raw
}
