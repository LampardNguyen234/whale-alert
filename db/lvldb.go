package db

import (
	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb"
)

type LevelDB struct {
	*leveldb.DB
}

// NewLvlDB returns a new LevelDB instance.
func NewLvlDB(path string) (*LevelDB, error) {
	ldb, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "levelDB.OpenFile fail")
	}
	return &LevelDB{DB: ldb}, nil
}

func (db *LevelDB) GetByKey(key []byte) ([]byte, error) {
	return db.DB.Get(key, nil)
}

func (db *LevelDB) SetByKey(key []byte, value []byte) error {
	return db.DB.Put(key, value, nil)
}

func (db *LevelDB) Close() error {
	return db.DB.Close()
}
