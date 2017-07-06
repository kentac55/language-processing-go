package exercise

import "strings"

func P04(s string) map[string]int {
	res := make(map[string]int)
	order := []int{1, 5, 6, 7, 8, 9, 15, 16, 19}
	for i, v := range strings.Split(s, " ") {
		i++
		for _, o := range order {
			if i == o {
				res[v[0:1]] = i
			} else {
				res[v[0:2]] = i
			}
		}
	}
	return res
}
