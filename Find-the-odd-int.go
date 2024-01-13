package kata

func FindOdd(seq []int) int {
	countMap := make(map[int]int)

	for _, num := range seq {
		countMap[num]++
	}

	for num, count := range countMap {
		if count%2 != 0 {
			return num
		}
	}
	return -1
}
