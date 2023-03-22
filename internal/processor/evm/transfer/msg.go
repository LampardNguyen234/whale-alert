package evm_transfer

import "fmt"

type Msg struct {
	From      string
	To        string
	Amount    string
	Token     string
	TokenName string
	TxHash    string
}

func (msg Msg) String() string {
	return fmt.Sprintf("%v %vs (%v) has just been transferred from %v to %v. TxHash: %v",
		msg.Amount,
		msg.TokenName,
		msg.Token,
		msg.From,
		msg.To,
		msg.TxHash,
	)
}
