package system

import (
	"crypto/rand"
	"math/big"
)

func GenerateLoginCode(userID string) (string, error) {
	code, err := rand.Int(rand.Reader, big.NewInt(10_000))
	if err != nil {
		return "", err
	}
	return code.String(), nil
}
