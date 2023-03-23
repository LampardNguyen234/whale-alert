package account_watcher

import (
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
)

type Direction string

const (
	directionIn   Direction = "IN"
	directionOut  Direction = "OUT"
	directionSelf Direction = "SELF"
)

type Msg struct {
	processorCommon.TxMsg
	Account   string
	Direction Direction
	Type      string
}

func (msg Msg) String() string {
	msgFormatter := new(processorCommon.MsgFormatter).
		FormatTitle("Account Watching").
		FormatKeyValueMsg("Account", msg.Account).
		FormatKeyValueMsg("Direction", msg.Direction).
		FormatKeyValueMsg("Amount", msg.Amount).
		FormatKeyValueMsg("Type", msg.Type).
		FormatKeyValueMsg("TxHash", processorCommon.FormatTxURL(msg.TxHash))

	return msgFormatter.String()
}
