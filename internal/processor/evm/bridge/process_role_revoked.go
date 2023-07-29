package bridge

import (
	"context"
	"fmt"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
	"github.com/LampardNguyen234/whale-alert/internal/processor/evm/bridge/contracts"
)

type RoleRevokedMsg struct {
	Network        string
	Operator       string
	RevokedAddress string
	Role           string
	TxHash         string
}

func (msg RoleRevokedMsg) String() string {
	msgFormatter := new(processorCommon.MsgFormatter).
		FormatTitle("Bridge Role Revoked").
		FormatKeyValueMsg("Operator", msg.Operator).
		FormatKeyValueMsg("RevokedAddress", msg.RevokedAddress).
		FormatKeyValueMsg("Role", msg.Role).
		FormatKeyValueMsg("Network", msg.Network).
		FormatKeyValueMsg("TxHash", msg.TxHash)

	return msgFormatter.String()
}

func (p *BridgeProcessor) processRoleRevoked(_ context.Context, eventMsg EventMsg) error {
	e := eventMsg.Event.(*contracts.BridgeRoleRevoked)
	role := p.suites[eventMsg.DomainID].Roles()[e.Role]
	if role == "" {
		role = fmt.Sprintf("%x", e.Role)
	}

	msg := RoleRevokedMsg{
		Operator:       e.Sender.String(),
		RevokedAddress: e.Account.String(),
		Role:           role,
		TxHash:         p.suites[eventMsg.DomainID].FormatTxHashLink(eventMsg.Event.GetLog().TxHash.String()),
		Network:        p.formatNetwork(eventMsg.DomainID),
	}

	return p.Whm.Alert(msg.String())
}
