package cosmos

import (
	"context"
	"testing"
)

func TestCosmosClient_GetValidator(t *testing.T) {
	resp, err := c.GetValidatorDetail(context.Background(), "astravaloper1q4k9pzj9srv6lnu2xsf4ctaw9wzuxf6sck60k2")
	if err != nil {
		panic(err)
	}

	c.log.Infof("%v", *resp)
}
