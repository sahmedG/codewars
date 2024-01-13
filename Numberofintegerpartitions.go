package kata

func Partitions(n int) int {
	parts := make([][]int, n+1)
	for i := range parts {
		parts[i] = make([]int, n+1)
	}
	for i := 0; i <= n; i++ {
		parts[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j >= i {
				parts[i][j] = parts[i-1][j] + parts[i][j-i]
			} else {
				parts[i][j] = parts[i-1][j]
			}
		}
	}
	return parts[n][n]
}
