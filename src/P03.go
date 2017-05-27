package exercise

import "strings"

func P03(s string) []int {
	res := []int{}
	for _, v := range strings.Split(s, " ") {
		res = append(res, len(v))
	}
	return res
}
