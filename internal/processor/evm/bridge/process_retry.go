package bridge

import (
	"context"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
	"github.com/LampardNguyen234/whale-alert/internal/processor/evm/bridge/contracts"
)

type RetryMsg struct {
	Network   string
	TxToRetry string
	TxHash    string
}

func (msg RetryMsg) String() string {
	msgFormatter := new(processorCommon.MsgFormatter).
		FormatTitle("Bridge Retry").
		FormatKeyValueMsg("TxToRetry", msg.TxToRetry).
		FormatKeyValueMsg("Network", msg.Network).
		FormatKeyValueMsg("TxHash", msg.TxHash)

	return msgFormatter.String()
}

func (p *BridgeProcessor) processRetry(_ context.Context, eventMsg EventMsg) error {
	msg := RetryMsg{
		TxToRetry: eventMsg.Event.(*contracts.BridgeRetry).TxHash,
		TxHash:    p.suites[eventMsg.DomainID].FormatTxHashLink(eventMsg.Event.GetLog().TxHash.String()),
		Network:   p.formatNetwork(eventMsg.DomainID),
	}

	return p.Whm.Alert(msg.String())
}
