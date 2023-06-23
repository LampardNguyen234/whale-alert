package volume_watch

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/clients/tiki"
	"github.com/LampardNguyen234/whale-alert/internal/common"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
	"time"
)

type Msg struct {
	tiki.Orders
	Period time.Duration
}

func (msg Msg) String() string {
	duration := "30 minutes"
	if msg.Period > time.Hour {
		duration = fmt.Sprintf("%v hour(s)", msg.Period.Hours())
	} else {
		duration = fmt.Sprintf("%v minute(s)", msg.Period.Minutes())
	}
	msgFormatter := new(processorCommon.MsgFormatter).
		FormatTitle("Tiki Volume Watching").
		FormatMsg(fmt.Sprintf("%v ASAs have been traded with %v order(s) in the last %v. TotalSell: %v, TotalBuy: %v.",
			common.FormatAmount(msg.Orders.Amount()),
			len(msg.Orders),
			duration,
			common.FormatAmount(msg.Orders.AmountByType(tiki.OrderSellType)),
			common.FormatAmount(msg.Orders.AmountByType(tiki.OrderBuyType)),
		))

	return msgFormatter.String()
}
