package cosmos

import (
	"context"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/gogo/protobuf/grpc"
)

type StakingClient struct {
	stakingTypes.QueryClient
}

func NewStakingClient(conn grpc.ClientConn) StakingClient {
	return StakingClient{
		QueryClient: stakingTypes.NewQueryClient(conn),
	}
}

func (c *CosmosClient) GetValidatorDetail(ctx context.Context, valAddress string) (*stakingTypes.Validator, error) {
	resp, err := c.StakingClient.Validator(ctx, &stakingTypes.QueryValidatorRequest{ValidatorAddr: valAddress})
	if err != nil {
		return nil, err
	}

	return &resp.Validator, nil
}
