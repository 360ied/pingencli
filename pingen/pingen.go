package pingen

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

func GeneratePIN(digits int64) (string, error) {
	max := big.NewInt(10)
	max.Exp(max, big.NewInt(digits), nil)

	pinInt, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%0"+strconv.FormatInt(digits, 10)+"d", pinInt), nil
}
