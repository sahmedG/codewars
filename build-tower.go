package kata

import "strings"

func TowerBuilder(nFloors int) []string {
	var tower []string

	for i := 0; i < nFloors; i++ {
		spaces := strings.Repeat(" ", nFloors-i-1)
		stars := strings.Repeat("*", 2*i+1)
		row := spaces + stars + spaces
		tower = append(tower, row)
	}
	return tower
}
