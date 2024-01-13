package kata

import "math"

func RectangleRotation(a, b int) int {
	// your code here
	get := func(x int) int {
		return int(math.Floor(float64(x) / math.Sqrt2))
	}

	f := func(a, b int) int {
		return get(a)*get(b) + (get(a)+1)*(get(b)+1)
	}

	evenMin1 := func(x int) int {
		if x%2 == 0 {
			return x - 1
		}
		return x
	}

	return evenMin1(f(a, b))
}
