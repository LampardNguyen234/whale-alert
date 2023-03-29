package request

import "fmt"

// APIUpdateAccountDetail holds an account detail.
type APIUpdateAccountDetail struct {
	// Address is either an EVM or Cosmos address of the account.
	Address string `json:"address" binding:"required" validate:"required"`

	// Name is the identity of the account.
	Name string `json:"name" binding:"required" validate:"required"`

	// Monitored indicates whether the service should monitor this account.
	Monitored bool `json:"monitored"`
}

func (req *APIUpdateAccountDetail) IsValid() (bool, error) {
	if req.Address == "" {
		return false, fmt.Errorf("empty address")
	}

	if req.Name == "" {
		return false, fmt.Errorf("empty name")
	}

	return true, nil
}
