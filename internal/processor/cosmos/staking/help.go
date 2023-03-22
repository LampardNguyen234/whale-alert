package staking

import (
	"context"
	"fmt"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func (p *StakingProcessor) getValidatorName(ctx context.Context, valAddr string) string {
	validatorDetail, err := p.CosmosClient.GetValidatorDetail(ctx, valAddr)
	if err == nil {
		return parseValidatorDetail(valAddr, validatorDetail.Description)
	}

	return valAddr
}

func parseValidatorDetail(valAddr string, desc stakingTypes.Description) string {
	if desc.Moniker != "" {
		return fmt.Sprintf("%v (%v...%v)", desc.Moniker,
			valAddr[:12],
			valAddr[len(valAddr)-12:],
		)
	}

	return valAddr
}
