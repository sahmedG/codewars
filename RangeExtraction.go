package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func Solution(list []int) string {
	return strings.Join(groupConsecutives(list), ",")
}

func groupConsecutives(list []int) []string {
	var result []string

	for _, g := range groupByEnumerate(list) {
		r := make([]string, len(g))
		for i, x := range g {
			r[i] = strconv.Itoa(x)
		}
		if len(r) > 2 {
			result = append(result, fmt.Sprintf("%s-%s", r[0], r[len(r)-1]))
		} else {
			result = append(result, r...)
		}
	}

	return result
}

func groupByEnumerate(list []int) [][]int {
	var result [][]int
	var currentGroup []int

	for i := 0; i < len(list); i++ {
		if i > 0 && list[i-1]+1 != list[i] {
			result = append(result, currentGroup)
			currentGroup = []int{}
		}
		currentGroup = append(currentGroup, list[i])
	}

	if len(currentGroup) > 0 {
		result = append(result, currentGroup)
	}

	return result
}
