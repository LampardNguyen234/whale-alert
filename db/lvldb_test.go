package db

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/common"
	"github.com/magiconair/properties/assert"
	"testing"
)

var (
	lvl *LevelDB
)

func init() {
	var err error

	lvl, err = NewLvlDB(DefaultLevelDBConfig().Path)
	if err != nil {
		panic(err)
	}
}

func TestLevelDB_SetByKey(t *testing.T) {
	key := []byte{0}

	tmp, _ := lvl.GetByKey(key)
	fmt.Printf("tmp: %v\n", tmp)

	val := common.RandomHash().Bytes()

	err := lvl.SetByKey(key, val)
	if err != nil {
		panic(err)
	}

	v, err := lvl.GetByKey(key)
	if err != nil {
		panic(err)
	}

	fmt.Printf("got: %v, want: %v\n", v, val)

	assert.Equal(t, v, val)
}
