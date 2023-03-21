package store

import "github.com/LampardNguyen234/whale-alert/db"

type Store struct {
	db db.KeyValueReaderWriter
}

// NewStore creates a new Store with the given db.
func NewStore(db db.KeyValueReaderWriter) *Store {
	return &Store{db: db}
}
