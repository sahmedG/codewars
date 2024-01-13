package kata

import "strings"

func DecodeMorse(morseCode string) string {
	words := strings.Split(morseCode, "   ")
	var decoded string

	for _, word := range words {
		decodedWord := decodeWord(word)
		decoded += decodedWord + " "
	}

	return strings.TrimSpace(decoded)
}

func decodeWord(word string) string {
	chars := strings.Split(word, " ")
	var decodedWord string

	for _, char := range chars {
		if val, ok := MORSE_CODE[char]; ok {
			decodedWord += val
		}
	}

	return decodedWord
}
