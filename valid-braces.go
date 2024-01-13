package kata

func ValidBraces(str string) bool {
	stack := []rune{}

	bracesMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, char := range str {
		if opening, exists := bracesMap[char]; exists {
			// Closing brace
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return false
			}
			stack = stack[:len(stack)-1] // Pop from stack
		} else {
			// Opening brace
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}
