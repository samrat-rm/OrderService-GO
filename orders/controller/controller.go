package handler

import (
	"crypto/rand"
	"math/big"
)

const (
	idLength = 10
)

func GenerateOrderID() uint32 {
	maxValue := new(big.Int).SetUint64(^uint64(0))
	randomInt, _ := rand.Int(rand.Reader, maxValue)

	return uint32(randomInt.Uint64())
}
