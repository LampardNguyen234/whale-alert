package staking

import (
	"fmt"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
)

type DelegateMsg struct {
	processorCommon.TxMsg
	Validator string
}

func (msg DelegateMsg) String() string {
	msgFormatter := new(processorCommon.MsgFormatter).
		FormatTitle("Staking").
		FormatKeyValueMsg("Staker", msg.From).
		FormatKeyValueMsg("Validator", msg.Validator).
		FormatKeyValueMsg("Amount", msg.Amount).
		FormatKeyValueMsg("TxHash", processorCommon.FormatTxURL(msg.TxHash))

	return msgFormatter.String()
}

type UndelegateMsg struct {
	processorCommon.TxMsg
	Validator string
}

func (msg UndelegateMsg) String() string {
	msgFormatter := new(processorCommon.MsgFormatter).
		FormatTitle("Withdraw Staking").
		FormatKeyValueMsg("Staker", msg.From).
		FormatKeyValueMsg("From Validator", msg.Validator).
		FormatKeyValueMsg("Amount", msg.Amount).
		FormatKeyValueMsg("TxHash", processorCommon.FormatTxURL(msg.TxHash))

	return msgFormatter.String()
}

type CreateValidatorMsg struct {
	processorCommon.TxMsg
	Address    string
	Name       string
	Commission float64
}

func (msg CreateValidatorMsg) String() string {
	msgFormatter := new(processorCommon.MsgFormatter).
		FormatTitle("New Validator")
	if msg.Name != "" {
		msgFormatter = msgFormatter.FormatKeyValueMsg("Name", msg.Name)
	}
	msgFormatter = msgFormatter.
		FormatKeyValueMsg("Address", msg.Address).
		FormatKeyValueMsg("Commission", fmt.Sprintf("%v%%", msg.Commission*100)).
		FormatKeyValueMsg("TxHash", processorCommon.FormatTxURL(msg.TxHash))

	return msgFormatter.String()
}
