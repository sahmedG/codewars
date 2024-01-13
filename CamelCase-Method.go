package kata

import "strings"

func CamelCase(s string) string {
	// your code here
	words := strings.Fields(s)
	var camelCaseString string
	for _, word := range words {
		camelCaseString += strings.Title(word)
	}
	return camelCaseString
}
