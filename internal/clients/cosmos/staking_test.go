package cosmos

import (
	"context"
	"testing"
)

func TestCosmosClient_GetValidator(t *testing.T) {
	resp, err := c.GetValidatorDetail(context.Background(), "astravaloper1xx87hygrd2dd35qj6e8e0kyv7hre7gm9t9nqc3")
	if err != nil {
		panic(err)
	}

	c.log.Infof("%v", *resp)
}
