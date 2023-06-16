package common

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/sha3"
	"strings"
)

const HashLength = 32

type Hash [HashLength]byte

var ZeroHash = [HashLength]byte{0}

func NewHashFromString(hexHash string) (Hash, error) {
	hexHash = strings.Replace(hexHash, "0x", "", -1)
	if len(hexHash) != 2*HashLength {
		return ZeroHash, fmt.Errorf("expected string of length %v, got %v", 2*HashLength, len(hexHash))
	}
	tmpBytes, err := hex.DecodeString(hexHash)
	if err != nil {
		return ZeroHash, err
	}
	res := [32]byte{}
	copy(res[:], tmpBytes)

	return res, nil
}

func MustNewHashFromString(hexHash string) Hash {
	ret, err := NewHashFromString(hexHash)
	if err != nil {
		panic(err)
	}

	return ret
}

func Digest(values ...interface{}) Hash {
	toBeDigested := make([]byte, 0)
	for _, v := range values {
		if tmp, ok := v.([]byte); ok {
			toBeDigested = append(toBeDigested, tmp...)
		} else {
			toBeDigested = append(toBeDigested, []byte(fmt.Sprintf("%v", v))...)
		}
	}
	return sha3.Sum256(toBeDigested)
}

func RandomHash() Hash {
	b := make([]byte, HashLength)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	return sha3.Sum256(b)
}

func (h Hash) IsZero() bool {
	return h == ZeroHash
}

func (h Hash) String() string {
	return fmt.Sprintf("%06x", h[:])
}

func (h Hash) Bytes() []byte {
	return h[:]
}
