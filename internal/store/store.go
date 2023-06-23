package store

import (
	"github.com/LampardNguyen234/whale-alert/common"
	"github.com/LampardNguyen234/whale-alert/db"
)

type Store struct {
	db           db.KeyValueReaderWriter
	cachedTokens common.Cache
}

// NewStore creates a new Store with the given db.
func NewStore(db db.KeyValueReaderWriter) *Store {
	return &Store{db: db, cachedTokens: common.NewSimpleCache()}
}

func (s *Store) Init() error {
	allTokens, err := s.getAllTokenDetails()
	if err != nil || len(allTokens) == 0 {
		allTokens = defaultTokenDetails()
		for _, d := range allTokens {
			err = s.storeTokenDetail(d)
			if err != nil {
				return err
			}
		}
	}

	for addr, d := range allTokens {
		s.cachedTokens.SetDefault(addr, d)
	}

	return nil
}
