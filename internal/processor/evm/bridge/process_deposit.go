package bridge

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/common"
	processorCommon "github.com/LampardNguyen234/whale-alert/internal/processor/common"
	"github.com/LampardNguyen234/whale-alert/internal/processor/evm/bridge/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"
	"math/big"
)

type DepositMsg struct {
	From        string
	To          string
	FromNetwork string
	ToNetwork   string
	Token       string
	TokenName   string
	TokenSymbol string
	Amount      float64
	TxHash      string
}

func (msg DepositMsg) String() string {
	msgFormatter := new(processorCommon.MsgFormatter).
		FormatTitle("Bridge Transfer").
		FormatKeyValueMsg("From", msg.From).
		FormatKeyValueMsg("To", msg.To).
		FormatKeyValueMsg("From Network", msg.FromNetwork).
		FormatKeyValueMsg("To Network", msg.ToNetwork).
		FormatKeyValueMsg("Token", fmt.Sprintf("%v - %v (%v)", msg.TokenName, msg.TokenSymbol, msg.Token)).
		FormatKeyValueMsg("Amount", common.FormatAmount(msg.Amount)).
		FormatKeyValueMsg("TxHash", msg.TxHash)

	return msgFormatter.String()
}

func (p *BridgeProcessor) processDeposit(ctx context.Context, eventMsg EventMsg) error {
	msg, err := p.depositEventToMsg(ctx, eventMsg)
	if err != nil {
		return err
	}
	tokenDetail := p.Db.GetTokenDetail(msg.Token)
	if tokenDetail.TokenAddress != "" && tokenDetail.TokenName != "" {
		if msg.Amount < tokenDetail.WhaleDefinition {
			return nil
		}
	}

	return p.Whm.Alert(msg.String())
}

func (p *BridgeProcessor) depositEventToMsg(ctx context.Context, e EventMsg) (*DepositMsg, error) {
	deposit := e.Event.(*contracts.BridgeDeposit)
	jsb, _ := json.Marshal(deposit)
	p.Log.Debugf("new deposit: %v", string(jsb))

	fromSuite := p.suites[e.DomainID]
	fromNetwork := p.formatNetwork(e.DomainID)
	toNetwork := p.formatNetwork(deposit.DestinationDomainID)

	callOpts := &bind.CallOpts{
		Context: ctx,
	}

	erc20HandlerAddr, err := fromSuite.Bridge.ResourceIDToHandlerAddress(callOpts, deposit.ResourceID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve handler address for resource %v", deposit.ResourceID)
	}

	erc20Handler, err := contracts.NewErc20Handler(erc20HandlerAddr, fromSuite.Client)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse erc20Handler for addr %v", erc20HandlerAddr.String())
	}

	tokenAddress, err := erc20Handler.ResourceIDToTokenContractAddress(callOpts, deposit.ResourceID)
	if err != nil {
		return nil, err
	}

	tokenContract, err := contracts.NewErc20(tokenAddress, fromSuite.Client)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve token contract")
	}

	tokenName, err := tokenContract.Name(callOpts)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve token name")
	}

	tokenSymbol, err := tokenContract.Symbol(callOpts)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve token symbol")
	}

	tokenDecimals, err := tokenContract.Decimals(callOpts)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve token decimals")
	}

	depositAmount := new(big.Int).SetBytes(deposit.Data[:32])

	from := deposit.User.String()
	to := from
	if len(deposit.Data) >= 84 {
		recipientLength := new(big.Int).SetBytes(deposit.Data[32:64])
		to = fmt.Sprintf("%x", deposit.Data[64:64+recipientLength.Int64()])
	}

	return &DepositMsg{
		From:        from,
		To:          to,
		FromNetwork: fromNetwork,
		ToNetwork:   toNetwork,
		Token:       tokenAddress.String(),
		TokenName:   tokenName,
		TokenSymbol: tokenSymbol,
		Amount:      common.GetNormalizedValue(depositAmount, int(tokenDecimals)),
		TxHash:      fromSuite.FormatTxHashLink(deposit.GetLog().TxHash.String()),
	}, nil
}
