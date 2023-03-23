package store

import (
	"encoding/base64"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/libs/rand"
	"testing"
)

func TestStore_GetAccountDetail(t *testing.T) {
	for i := 0; i < numTests; i++ {
		detail := randomAccountDetail()

		err := s.StoreAccountDetail(detail)
		if err != nil {
			panic(err)
		}

		tmpDetail, err := s.GetAccountDetail(detail.Address)
		if err != nil {
			panic(err)
		}

		assert.Equal(t, compareAccountDetail(*tmpDetail, detail), true, "%v accounts mismatch: %v vs %v", i, tmpDetail, detail)
	}
}

func TestStore_GetAllAccountDetails(t *testing.T) {
	currentNumAccounts := 0
	allAccounts, err := s.GetAllAccountDetails()
	if err == nil {
		currentNumAccounts = len(allAccounts)
	}

	numAccounts := 1 + rand.Int()%numTests
	for i := 0; i < numAccounts; i++ {
		err = s.StoreAccountDetail(randomAccountDetail())
		if err != nil {
			panic(err)
		}
	}

	allAccounts, err = s.GetAllAccountDetails()
	if err != nil {
		panic(err)
	}

	assert.Equal(t, len(allAccounts)-currentNumAccounts, numAccounts, "expect %v accounts, got %v", numAccounts, len(allAccounts)-currentNumAccounts)
}

func TestStore_IsAccountMonitored(t *testing.T) {
	for i := 0; i < numTests; i++ {
		detail := randomAccountDetail()
		err := s.StoreAccountDetail(detail)
		if err != nil {
			panic(err)
		}

		isActuallyMonitored, _ := s.IsAccountMonitored(detail.Address)

		assert.Equal(t, isActuallyMonitored, detail.Monitored, "expect monitored to be %v, got %v", i, detail.Monitored, isActuallyMonitored)
	}
}

func TestStore_GetAllMonitoredAccounts(t *testing.T) {
	numAccounts := 1 + rand.Int()%numTests
	newMonitoredAccounts := make(map[string]*AccountDetail)
	for i := 0; i < numAccounts; i++ {
		d := randomAccountDetail()
		err := s.StoreAccountDetail(d)
		if err != nil {
			panic(err)
		}
		if d.Monitored {
			newMonitoredAccounts[d.Address] = &d
		}
	}

	allAccounts, err := s.GetAllMonitoredAccounts()
	if err != nil {
		panic(err)
	}

	for addr, detail := range newMonitoredAccounts {
		if tmpDetail, ok := allAccounts[addr]; !ok {
			panic(fmt.Sprintf("expected account %v to exist", addr))
		} else {
			assert.Equal(t, compareAccountDetail(*tmpDetail, *detail), true, "accounts mismatch: %v vs %v", tmpDetail, detail)
		}
	}
}

func randomAccountDetail() AccountDetail {
	return AccountDetail{
		Address:   common.BytesToAddress(rand.Bytes(32)).String(),
		Name:      base64.StdEncoding.EncodeToString(rand.Bytes(10)),
		Monitored: rand.Int()%2 == 1,
	}
}

func compareAccountDetail(a, b AccountDetail) bool {
	if a.Address != b.Address {
		return false
	}

	if a.Name != b.Name {
		return false
	}

	return true
}
