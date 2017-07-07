package exercise

import "strings"

func P04(s string) map[string]int {
	res := make(map[string]int, 20)
	//order := []int{1, 5, 6, 7, 8, 9, 15, 16, 19}
	for i, v := range strings.Split(s, " ") {
		i++
		if i == 1 || i == 5 || i == 6 || i == 7 || i == 8 || i == 9 || i == 15 || i == 16 || i == 19 {
			res[v[0:1]] = i
		} else {
			res[v[0:2]] = i
		}
	}
	return res
}
