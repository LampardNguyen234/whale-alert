package cosmos

import (
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"google.golang.org/grpc"
)

type BankClient struct {
	bankTypes.QueryClient
	bankTypes.MsgClient
}

func NewBankClient(conn *grpc.ClientConn) BankClient {
	return BankClient{
		QueryClient: bankTypes.NewQueryClient(conn),
		MsgClient:   bankTypes.NewMsgClient(conn),
	}
}
