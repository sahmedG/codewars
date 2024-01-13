package kata

func IsValidWalk(walk []rune) bool {
	if len(walk) != 10 {
		return false
	}

	x, y := 0, 0
	for _, direction := range walk {
		switch direction {
		case 'n':
			y++
		case 's':
			y--
		case 'e':
			x++
		case 'w':
			x--
		}
	}

	return x == 0 && y == 0
}
