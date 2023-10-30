package staking

import (
	"context"
	"fmt"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func (p *StakingProcessor) getValidatorName(_ context.Context, valAddr string) string {
	validatorDetail, err := p.CosmosClient.GetValidatorDetail(valAddr)
	if err == nil {
		return parseValidatorDetail(valAddr, validatorDetail.Description)
	}

	return valAddr
}

func parseValidatorDetail(valAddr string, desc stakingTypes.Description) string {
	if desc.Moniker != "" {
		return fmt.Sprintf("%v (%v...%v)", desc.Moniker,
			valAddr[:15],
			valAddr[len(valAddr)-15:],
		)
	}

	return valAddr
}
