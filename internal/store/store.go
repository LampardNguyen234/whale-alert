package store

import (
	"github.com/LampardNguyen234/whale-alert/db"
	"sync"
)

type Store struct {
	db        db.KeyValueReaderWriter
	allTokens map[string]TokenDetail
	mtx       *sync.Mutex
}

// NewStore creates a new Store with the given db.
func NewStore(db db.KeyValueReaderWriter) *Store {
	return &Store{db: db, allTokens: make(map[string]TokenDetail), mtx: new(sync.Mutex)}
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

	s.allTokens = allTokens
	return nil
}
