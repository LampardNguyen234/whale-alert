package bridge

import (
	"context"
	"fmt"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
	"github.com/LampardNguyen234/whale-alert/internal/processor/evm/bridge/contracts"
)

type RoleGrantedMsg struct {
	Network  string
	Operator string
	Grantee  string
	Role     string
	TxHash   string
}

func (msg RoleGrantedMsg) String() string {
	msgFormatter := new(processorCommon.MsgFormatter).
		FormatTitle("Bridge Role Granted").
		FormatKeyValueMsg("Operator", msg.Operator).
		FormatKeyValueMsg("RevokedAddress", msg.Grantee).
		FormatKeyValueMsg("Role", msg.Role).
		FormatKeyValueMsg("Network", msg.Network).
		FormatKeyValueMsg("TxHash", msg.TxHash)

	return msgFormatter.String()
}

func (p *BridgeProcessor) processRoleGranted(_ context.Context, eventMsg EventMsg) error {
	e := eventMsg.Event.(*contracts.BridgeRoleGranted)
	role := p.suites[eventMsg.DomainID].Roles()[e.Role]
	if role == "" {
		role = fmt.Sprintf("%x", e.Role)
	}

	msg := RoleGrantedMsg{
		Operator: e.Sender.String(),
		Grantee:  e.Account.String(),
		Role:     role,
		TxHash:   p.suites[eventMsg.DomainID].FormatTxHashLink(eventMsg.Event.GetLog().TxHash.String()),
		Network:  p.formatNetwork(eventMsg.DomainID),
	}

	return p.Whm.Alert(msg.String())
}
