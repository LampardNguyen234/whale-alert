package store

import (
	"encoding/json"
	"fmt"
	"github.com/LampardNguyen234/whale-alert/common"
	"github.com/LampardNguyen234/whale-alert/db"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type AccountDetail struct {
	Address string `json:"Address"`

	Name string `json:"Name"`

	Monitored bool `json:"Monitored,omitempty"`
}

func (d AccountDetail) String() string {
	if d.Name == "" {
		return fmt.Sprintf("%v", d.Address)
	}

	return fmt.Sprintf("%v (%v...%v)", d.Name, d.Address[:10], d.Address[len(d.Address)-10:])
}

// StoreAccountDetail stores the given AccountDetail to db.
func (s *Store) StoreAccountDetail(detail AccountDetail) error {
	addr, err := common.AccountAddressToEthAddr(detail.Address)
	if err != nil {
		return err
	}

	jsb, err := json.Marshal(detail)
	if err != nil {
		return err
	}
	return s.db.SetByKey(makeKey(addressDetailKey, addr.Bytes()...), jsb)
}

// GetAccountDetail retrieves the detail of an Account given its address.
func (s *Store) GetAccountDetail(address string) (*AccountDetail, error) {
	addr, err := common.AccountAddressToEthAddr(address)
	if err != nil {
		return nil, err
	}

	data, err := s.db.GetByKey(makeKey(addressDetailKey, addr.Bytes()...))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("no account found")
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

// IsAccountMonitored checks if the given address is monitored.
func (s *Store) IsAccountMonitored(address string) (bool, error) {
	account, err := s.GetAccountDetail(address)
	if err != nil {
		return false, err
	}

	return account.Monitored, nil
}

// GetAllMonitoredAccounts retrieves all monitored account.
func (s *Store) GetAllMonitoredAccounts() (map[string]*AccountDetail, error) {
	allAccounts, err := s.GetAllAccountDetails()
	if err != nil {
		return nil, err
	}

	ret := make(map[string]*AccountDetail)
	for _, account := range allAccounts {
		if account.Monitored {
			ret[account.Address] = &AccountDetail{
				Address:   account.Address,
				Name:      account.Name,
				Monitored: account.Monitored,
			}
		}
	}

	return ret, nil
}
