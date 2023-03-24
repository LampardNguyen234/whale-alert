package request

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/store"
)

// APIUpdateTokenDetail holds a token detail.
type APIUpdateTokenDetail store.TokenDetail

func (req *APIUpdateTokenDetail) IsValid() (bool, error) {
	if req.TokenAddress == "" {
		return false, fmt.Errorf("empty address")
	}
	if req.TokenName == "" {
		req.TokenName = req.TokenAddress
	}
	if req.Decimals == 0 {
		return false, fmt.Errorf("invalid decimal")
	}

	return true, nil
}
