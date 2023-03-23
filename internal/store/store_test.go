package store

import "github.com/LampardNguyen234/whale-alert/db"

var (
	s        *Store
	numTests = 100
)

func init() {
	lvlDB, err := db.NewLvlDB("./data")
	if err != nil {
		panic(err)
	}

	s = NewStore(lvlDB)
}
