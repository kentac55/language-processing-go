package exercise

import (
	"regexp"
	"strings"
)

func P03(s string) []int {
	res := []int{}
	reg := regexp.MustCompile(`[^0-9A-Za-z ]`)
	s = reg.ReplaceAllString(s, "")
	for _, v := range strings.Split(s, " ") {
		res = append(res, len(v))
	}
	return res
}
