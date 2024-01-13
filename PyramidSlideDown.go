package kata

func LongestSlideDown(pyramid [][]int) int {
	// your code here
	for row := len(pyramid) - 2; row >= 0; row-- {
		for i := 0; i < len(pyramid[row]); i++ {
			pyramid[row][i] += max(pyramid[row+1][i], pyramid[row+1][i+1])
		}
	}

	return pyramid[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
