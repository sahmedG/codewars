package kata

import "math"

func ProperFractions(n int) int {
	// your code here
	if n == 1 {
		return 0
	}

	a := 1
	b := n
	c := int(math.Sqrt(float64(n))) + 1

	for i := 2; i <= c; i++ {
		if mine(i) {
			if b%i == 0 {
				a *= i - 1
				b /= i
			}
			for b%i == 0 {
				a *= i
				b /= i
			}
		}
	}

	if b > 1 {
		a *= b - 1
	}

	return a
}

func mine(z int) bool {
	if z <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(z))); i++ {
		if z%i == 0 {
			return false
		}
	}
	return true
}
