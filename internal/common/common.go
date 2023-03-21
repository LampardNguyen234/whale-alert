package common

import (
	"math/big"
	"time"
)

const (
	AsaDecimals      = 18
	DefaultSleepTime = 2 * time.Second
	AsaAddress       = "0x0000000000000000000000000000000000000000"
)

var AsaDecimalsBigInt = big.NewInt(AsaDecimals)
