package transfer

import (
	"fmt"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
)

type Msg struct {
	From      string
	To        string
	Amount    string
	Token     string
	TokenName string
	TxHash    string
}

func (msg Msg) String() string {
	msgFormatter := new(processorCommon.MsgFormatter).
		FormatTitle("Cosmos Transfer").
		FormatKeyValueMsg("From", msg.From).
		FormatKeyValueMsg("To", msg.To).
		FormatKeyValueMsg("Amount", msg.Amount).
		FormatKeyValueMsg("Token", fmt.Sprintf("%v (%v)", msg.TokenName, msg.Token)).
		FormatKeyValueMsg("TxHash", processorCommon.FormatTxURL(msg.TxHash))

	return msgFormatter.String()
}
