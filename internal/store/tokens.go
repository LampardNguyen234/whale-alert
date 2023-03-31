package store

import (
	"encoding/json"
	"fmt"
	"github.com/LampardNguyen234/whale-alert/db"
	"github.com/LampardNguyen234/whale-alert/internal/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/syndtr/goleveldb/leveldb/util"
	"strings"
)

type TokenDetail struct {
	// TokenName is the name of the token.
	TokenName string `json:"TokenName,omitempty" binding:"required"`

	// TokenAddress is the EVM address of the token.
	TokenAddress string `json:"TokenAddress" binding:"required"`

	// Decimals is the number of decimal places of the token.
	Decimals int `json:"Decimals" binding:"required"`

	// WhaleDefinition is the amount to trigger the whale alerts.
	WhaleDefinition float64 `json:"WhaleDefinition,omitempty"`
}

func (s *Store) UpdateTokenDetail(d TokenDetail) error {
	err := s.storeTokenDetail(d)
	if err != nil {
		return err
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.allTokens[d.TokenAddress] = d

	return nil
}

func (s *Store) GetTokenDetail(tokenAddress string) TokenDetail {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return s.allTokens[tokenAddress]
}

func (s *Store) GetAllTokenDetails() map[string]TokenDetail {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	ret := make(map[string]TokenDetail)
	for token, d := range s.allTokens {
		ret[token] = d
	}

	return ret
}

// storeTokenDetail stores the detail of a token
func (s *Store) storeTokenDetail(d TokenDetail) error {
	addressBytes := tokenAddressToBytes(d.TokenAddress)
	jsb, err := json.Marshal(d)
	if err != nil {
		return err
	}

	return s.db.SetByKey(makeKey(tokenDetailKey, addressBytes...), jsb)
}

// getTokenDetail retrieves the detail of a token.
func (s *Store) getTokenDetail(tokenAddress string) (*TokenDetail, error) {
	addressBytes := tokenAddressToBytes(tokenAddress)
	data, err := s.db.GetByKey(makeKey(tokenDetailKey, addressBytes...))
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

// getAllTokenDetails retrieves all saved tokens details.
func (s *Store) getAllTokenDetails() (map[string]TokenDetail, error) {
	levelDB, ok := s.db.(*db.LevelDB)
	if !ok {
		return nil, fmt.Errorf("method not supported")
	}
	iter := levelDB.DB.NewIterator(util.BytesPrefix(makePrefix(addressDetailKey)), nil)
	res := make(map[string]TokenDetail)
	for iter.Next() {
		var d TokenDetail
		err := json.Unmarshal(iter.Value(), &d)
		if err != nil {
			return nil, err
		}
		res[d.TokenAddress] = d
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		return nil, fmt.Errorf("iteration error: %v", err)
	}

	return res, nil
}

func defaultTokenDetails() map[string]TokenDetail {
	return map[string]TokenDetail{
		common.AsaAddress: {
			TokenName:       "ASA",
			TokenAddress:    common.AsaAddress,
			Decimals:        common.AsaDecimals,
			WhaleDefinition: 10000,
		},
	}
}

func tokenAddressToBytes(address string) []byte {
	address = strings.ToLower(address)
	address = strings.Replace(address, "0x", "", -1)
	if address == "" {
		address = common.AsaAddress
	}

	return crypto.Keccak256([]byte(address))[:]
}
