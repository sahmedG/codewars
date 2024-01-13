package kata

import "math"

func ListSquared(m, n int) [][]int {
	// your code
	res := [][]int{}
	for num := m; num <= n; num++ {
		divs := getDivs(num)
		divsSum := sumSqr(divs)

		if isPerSqr(divsSum) {
			res = append(res, []int{num, divsSum})
		}
	}
	return res
}

func getDivs(num int) []int {
	divs := []int{}
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			divs = append(divs, i*i)
		}
	}

	return divs
}

func sumSqr(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func isPerSqr(num int) bool {
	root := int(math.Sqrt(float64(num)))
	return root*root == num
}
