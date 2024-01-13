package kata

func FindEvenIndex(arr []int) int {
	for i := 0; i < len(arr); i++ {
		leftsum := sum(arr[:i])
		rightsum := sum(arr[i+1:])

		if leftsum == rightsum {
			return i
		}
	}
	return -1
}

func sum(slice []int) int {
	tot := 0
	for _, num := range slice {
		tot += num
	}
	return tot
}
