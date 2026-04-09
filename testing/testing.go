package testing

import (
	"crypto/rand"
	"math/big"
)

func RandomRangeValue(min, max int) int {
	a, _ := rand.Int(rand.Reader, big.NewInt(100))
	return int(a.Int64())%(max-min+1) + min
}
