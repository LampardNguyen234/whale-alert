package store

import (
	"encoding/json"
	"fmt"
)

type TokenDetail struct {
	TokenName    string `json:"TokenName"`
	TokenAddress string `json:"TokenAddress"`
	Decimals     int    `json:"Decimals"`
}

// StoreTokenDetail stores the detail of a token
func (s *Store) StoreTokenDetail(d TokenDetail) error {
	jsb, err := json.Marshal(d)
	if err != nil {
		return err
	}

	return s.db.SetByKey(makeKey(tokenDetailKey, []byte(d.TokenAddress)...), jsb)
}

// GetTokenDetail retrieves the detail of a token.
func (s *Store) GetTokenDetail(tokenAddress string) (*TokenDetail, error) {
	data, err := s.db.GetByKey(makeKey(tokenDetailKey, []byte(tokenAddress)...))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("empty data")
	}

	var ret TokenDetail
	err = json.Unmarshal(data, &ret)
	if err != nil {
		return nil, fmt.Errorf("failed to decode db data: %v", err)
	}

	return &ret, nil
}
