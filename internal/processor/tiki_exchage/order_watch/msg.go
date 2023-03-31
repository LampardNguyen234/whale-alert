package order_watch

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/clients/tiki"
	"github.com/LampardNguyen234/whale-alert/internal/common"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
)

type Direction string

type Msg struct {
	tiki.Order
	CurrentPrice float64
}

func (msg Msg) String() string {
	msgFormatter := new(processorCommon.MsgFormatter).
		FormatTitle("Tiki Order Watching").
		FormatKeyValueMsg("Amount", fmt.Sprintf("%v (%v VND)",
			common.FormatAmount(msg.Amount), common.FormatAmount(msg.Total))).
		FormatKeyValueMsg("Price/LastPrice", fmt.Sprintf("%v/%v",
			common.FormatAmount(msg.Price), common.FormatAmount(msg.CurrentPrice))).
		FormatKeyValueMsg("Side", msg.Side).
		FormatKeyValueMsg("ID", msg.ID).
		FormatKeyValueMsg("Time", msg.CreatedAt.String())

	return msgFormatter.String()
}
