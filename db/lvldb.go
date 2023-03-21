package db

import (
	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb"
)

type LevelDB struct {
	db *leveldb.DB
}

// NewLvlDB returns a new LevelDB instance.
func NewLvlDB(path string) (*LevelDB, error) {
	ldb, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "levelDB.OpenFile fail")
	}
	return &LevelDB{db: ldb}, nil
}

func (db *LevelDB) GetByKey(key []byte) ([]byte, error) {
	return db.db.Get(key, nil)
}

func (db *LevelDB) SetByKey(key []byte, value []byte) error {
	return db.db.Put(key, value, nil)
}

func (db *LevelDB) Close() error {
	return db.db.Close()
}
