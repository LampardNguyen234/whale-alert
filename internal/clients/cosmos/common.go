package cosmos

import (
	"github.com/LampardNguyen234/whale-alert/internal/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	vestingTypes "github.com/evmos/evmos/v6/x/vesting/types"
	"math/big"
)

const (
	Denom = "aastra"
)

// ParseCosmosMsgValue returns the value of the given Cosmos message.
func (c *CosmosClient) ParseCosmosMsgValue(msg sdk.Msg) float64 {
	v := big.NewInt(0)
	switch msg.(type) {
	case *bankTypes.MsgSend:
		v = msg.(*bankTypes.MsgSend).Amount.AmountOf(Denom).BigInt()
	case *bankTypes.MsgMultiSend:
		tmpMsg := msg.(*bankTypes.MsgMultiSend)
		for _, out := range tmpMsg.Outputs {
			v = v.Add(v, out.Coins.AmountOf(Denom).BigInt())
		}

	case *stakingTypes.MsgDelegate:
		v = msg.(*stakingTypes.MsgDelegate).Amount.Amount.BigInt()
	case *stakingTypes.MsgBeginRedelegate:
		v = msg.(*stakingTypes.MsgBeginRedelegate).Amount.Amount.BigInt()
	case *stakingTypes.MsgUndelegate:
		v = msg.(*stakingTypes.MsgUndelegate).Amount.Amount.BigInt()
	case *stakingTypes.MsgCreateValidator:
		v = msg.(*stakingTypes.MsgCreateValidator).Value.Amount.BigInt()

	case *vestingTypes.MsgCreateClawbackVestingAccount:
		tmpMsg := msg.(*vestingTypes.MsgCreateClawbackVestingAccount)
		for _, period := range tmpMsg.VestingPeriods {
			v = v.Add(v, period.Amount.AmountOf(Denom).BigInt())
		}

	default:
		v = big.NewInt(0)
	}

	return common.GetNormalizedValue(v)
}

// ParseCosmosMsgSender returns the sender of the given Cosmos message.
func (c *CosmosClient) ParseCosmosMsgSender(msg sdk.Msg) string {
	ret := ""
	switch msg.(type) {
	case *bankTypes.MsgSend,
		*stakingTypes.MsgCreateValidator,
		*stakingTypes.MsgDelegate,
		*stakingTypes.MsgBeginRedelegate,
		*stakingTypes.MsgUndelegate,
		*vestingTypes.MsgCreateClawbackVestingAccount,
		*vestingTypes.MsgClawback:
		ret = msg.GetSigners()[0].String()
	default:
	}

	return ret
}

// ParseCosmosMsgReceiver returns the recipients of the given Cosmos message.
func (c *CosmosClient) ParseCosmosMsgReceiver(msg sdk.Msg) string {
	ret := ""
	switch msg.(type) {
	case *bankTypes.MsgSend:
		ret = msg.(*bankTypes.MsgSend).ToAddress
	case *vestingTypes.MsgCreateClawbackVestingAccount:
		ret = msg.(*vestingTypes.MsgCreateClawbackVestingAccount).ToAddress
	case *stakingTypes.MsgCreateValidator,
		*stakingTypes.MsgDelegate,
		*stakingTypes.MsgBeginRedelegate,
		*stakingTypes.MsgUndelegate,
		*vestingTypes.MsgClawback:
	default:
	}

	return ret
}
