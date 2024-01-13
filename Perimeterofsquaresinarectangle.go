package kata

func Perimeter(n int) int {
	// your code
	fibseq := []int{0, 1}
	for i := 2; i <= n+1; i++ {
		fibseq = append(fibseq, fibseq[i-1]+fibseq[i-2])
	}
	totper := 4 * sums(fibseq)
	return totper
}

func sums(nums []int) int {
	res := 0
	for _, num := range nums {
		res += num
	}
	return res
}
