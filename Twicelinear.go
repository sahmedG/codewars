package kata

func DblLinear(n int) int {
	// your code
	u := []int{1}
	x, y := 0, 0

	for len(u) <= n {
		nx := 2*u[x] + 1
		ny := 3*u[y] + 1

		if nx < ny {
			u = append(u, nx)
			x++
		} else if nx > ny {
			u = append(u, ny)
			y++
		} else {
			u = append(u, nx)
			x++
			y++
		}
	}
	return u[n]
}
