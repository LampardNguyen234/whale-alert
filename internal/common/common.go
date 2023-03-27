package common

import (
	"github.com/dustin/go-humanize"
	"math/big"
	"time"
)

const (
	AsaDecimals      = 18
	DefaultSleepTime = 2 * time.Second
	AsaAddress       = "0x"
	DecimalDigits    = 5
)

var AsaDecimalsBigInt = big.NewInt(AsaDecimals)

func GetNormalizedValue(rawValue *big.Int, decimals ...int) float64 {
	if rawValue == nil {
		return 0
	}

	decimalBigInt := AsaDecimalsBigInt
	if len(decimals) > 0 {
		decimalBigInt = new(big.Int).SetInt64(int64(decimals[0]))
	}

	amt := new(big.Float).SetInt(rawValue)
	amt = amt.Quo(amt, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), decimalBigInt, nil)))
	amtFloat, _ := amt.Float64()

	return amtFloat
}

func FormatAmount(amt float64) string {
	return humanize.FormatFloat("#,###.##", amt)
}
