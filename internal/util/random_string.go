package util

import (
	"crypto/rand"
	"math/big"
)

func RandomString(n int) (string, error) {
	runes := []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-")
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(runes))))
		if err != nil {
			return "", err
		}
		ret[i] = byte(runes[num.Int64()])
	}

	return string(ret), nil
}
