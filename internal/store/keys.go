package store

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	keyPrefixLength = 12
)

const (
	lastBlkKey = iota
	processedTxKey
	tokenDetailKey
)

func makePrefix(keyPrefix int) []byte {
	res := crypto.Keccak256([]byte(fmt.Sprintf("%v", keyPrefix)))
	res = append(res[:keyPrefixLength], make([]byte, 32-keyPrefixLength)...)
	return res
}

func makeKey(keyPrefix int, aux ...byte) []byte {
	res := makePrefix(keyPrefix)
	if len(aux) > 0 {
		res = append(res[:keyPrefixLength], crypto.Keccak256(aux)[:32-keyPrefixLength]...)
	}
	return res
}
