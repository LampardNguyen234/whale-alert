package store

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/LampardNguyen234/whale-alert/db"
	"github.com/syndtr/goleveldb/leveldb/util"
	"strings"
)

type AccountDetail struct {
	Address string `json:"Address"`

	Name string `json:"Name"`
}

// StoreAccountDetail stores the given AccountDetail to db.
func (s *Store) StoreAccountDetail(detail AccountDetail) error {
	address := strings.Replace(detail.Address, "0x", "", -1)
	address = strings.Replace(address, "0X", "", -1)
	addressBytes, err := hex.DecodeString(address)
	if err != nil {
		return err
	}

	jsb, err := json.Marshal(detail)
	if err != nil {
		return err
	}
	return s.db.SetByKey(makeKey(addressDetailKey, addressBytes...), jsb)
}

// GetAccountDetail retrieves the detail of an Account given its address.
func (s *Store) GetAccountDetail(address string) (*AccountDetail, error) {
	address = strings.Replace(address, "0x", "", -1)
	address = strings.Replace(address, "0X", "", -1)
	addressBytes, err := hex.DecodeString(address)
	if err != nil {
		return nil, err
	}

	data, err := s.db.GetByKey(makeKey(addressDetailKey, addressBytes...))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}

	var detail AccountDetail
	err = json.Unmarshal(data, &detail)
	if err != nil {
		return nil, err
	}

	return &detail, nil
}

// GetAllAccountDetails retrieves all saved account details.
func (s *Store) GetAllAccountDetails() ([]AccountDetail, error) {
	levelDB, ok := s.db.(*db.LevelDB)
	if !ok {
		return nil, fmt.Errorf("method not supported")
	}
	iter := levelDB.DB.NewIterator(util.BytesPrefix(makePrefix(addressDetailKey)), nil)
	res := make([]AccountDetail, 0)
	for iter.Next() {
		var d AccountDetail
		err := json.Unmarshal(iter.Value(), &d)
		if err != nil {
			return nil, err
		}
		res = append(res, d)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		return nil, fmt.Errorf("iteration error: %v", err)
	}

	return res, nil
}
