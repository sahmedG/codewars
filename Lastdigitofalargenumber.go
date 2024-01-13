package kata

import (
	"math/big"
)

func LastDigit(n1, n2 string) int {
	num1, _ := new(big.Int).SetString(n1, 10)
	num2, _ := new(big.Int).SetString(n2, 10)

	if num2.Cmp(big.NewInt(0)) == 0 {
		return 1
	}

	num1.Mod(num1, big.NewInt(10))
	num2.Mod(num2, big.NewInt(4))

	if num2.Cmp(big.NewInt(0)) == 0 {
		num2.Set(big.NewInt(4))
	}

	result := big.NewInt(1)
	for i := big.NewInt(0); i.Cmp(num2) < 0; i.Add(i, big.NewInt(1)) {
		result.Mul(result, num1)
		result.Mod(result, big.NewInt(10))
	}

	return int(result.Int64())

}
