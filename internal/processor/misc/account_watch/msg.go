package account_watcher

import (
	"fmt"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
)

type Direction string

const (
	directionIn      Direction = "IN"
	directionOut     Direction = "OUT"
	directionBothWay Direction = "BI"
)

type Msg struct {
	processorCommon.TxMsg
	Account   []string
	Direction Direction
	Type      string
}

func (msg Msg) String() string {
	accMsg := msg.Account[0]
	if len(msg.Account) > 1 {
		accMsg = fmt.Sprintf("%v  ==>  %v", msg.Account[0], msg.Account[1])
	}
	msgFormatter := new(processorCommon.MsgFormatter).
		FormatTitle("Account Watching").
		FormatKeyValueMsg("Account", accMsg).
		FormatKeyValueMsg("Direction", msg.Direction).
		FormatKeyValueMsg("Amount", msg.Amount).
		FormatKeyValueMsg("Type", msg.Type).
		FormatKeyValueMsg("TxHash", processorCommon.FormatTxURL(msg.TxHash))

	return msgFormatter.String()
}
