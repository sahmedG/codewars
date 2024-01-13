package kata

func Rot13(msg string) string {
	// Your code here
	str := ""
	for _, char := range msg {
		if char >= 65 && char < 78 || char >= 97 && char < 110 {
			char = char + 13
		} else if char >= 78 && char < 91 || char >= 110 && char < 123 {
			char = char - 13
		}
		str += string(char)
	}
	return str
}
