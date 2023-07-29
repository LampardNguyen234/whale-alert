package bridge

import (
	"context"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
	"github.com/LampardNguyen234/whale-alert/internal/processor/evm/bridge/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type RelayerRemovedMsg struct {
	Network       string
	Relayer       string
	TotalRelayers uint64
	TxHash        string
}

func (msg RelayerRemovedMsg) String() string {
	msgFormatter := new(processorCommon.MsgFormatter).
		FormatTitle("Bridge Relayer Removed").
		FormatKeyValueMsg("Address", msg.Relayer).
		FormatKeyValueMsg("Network", msg.Network)

	if msg.TotalRelayers != 0 {
		msgFormatter = msgFormatter.FormatKeyValueMsg("TotalRelayers", msg.TotalRelayers)
	}
	msgFormatter = msgFormatter.FormatKeyValueMsg("TxHash", msg.TxHash)

	return msgFormatter.String()
}

func (p *BridgeProcessor) processRelayerRemoved(_ context.Context, eventMsg EventMsg) error {
	msg := RelayerAddedMsg{
		Relayer: eventMsg.Event.(*contracts.BridgeRelayerRemoved).Relayer.String(),
		TxHash:  p.suites[eventMsg.DomainID].FormatTxHashLink(eventMsg.Event.GetLog().TxHash.String()),
		Network: p.formatNetwork(eventMsg.DomainID),
	}
	totalRelayers, err := p.suites[eventMsg.DomainID].Bridge.TotalRelayers(&bind.CallOpts{})
	if err != nil {
		p.Log.Errorf("failed to get total relayers: %v", err)
	} else {
		msg.TotalRelayers = totalRelayers.Uint64()
	}

	return p.Whm.Alert(msg.String())
}
