package kata

func Multiple3And5(number int) int {
	if number < 0 {
		return 0
	}
	sum := 0
	for i := 3; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
