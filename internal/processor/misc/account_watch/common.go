package account_watcher

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gogo/protobuf/proto"
	"strings"
)

const (
	EvmUnknownMsgType  = "EVM-UNKNOWN"
	EvmTransferMsgType = "EVM-TRANSFER"
)

func (p *AccountWatchProcessor) getEvmMsgType(receipt *types.Receipt, tx *types.Transaction) string {
	ret := EvmUnknownMsgType

	if len(receipt.Logs) == 0 {
		ret = EvmTransferMsgType
	}

	return ret
}

func (p *AccountWatchProcessor) getCosmosMsgType(msg sdk.Msg) string {
	msgType := proto.MessageName(msg)

	tmp := strings.Split(msgType, ".")
	return tmp[len(tmp)-1]
}
