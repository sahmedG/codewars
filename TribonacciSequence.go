package kata

func Tribonacci(signature [3]float64, n int) []float64 {
	if n == 0 {
		return []float64{}
	}

	tribonacciSeq := make([]float64, n)
	for i := 0; i < len(signature) && i < n; i++ {
		tribonacciSeq[i] = signature[i]
	}

	for i := 3; i < n; i++ {
		tribonacciSeq[i] = tribonacciSeq[i-1] + tribonacciSeq[i-2] + tribonacciSeq[i-3]
	}

	return tribonacciSeq
}
