package exercise

import (
	"math/rand"
	"strings"
)

func P09(s string, seed int64) string {
	rand.Seed(seed)
	res := ""
	for _, w := range strings.Split(s, " ") {
		if len(w) <= 4 {
			res += " " + w
			continue
		}
		r := []rune(w)
		l := make([]rune, 0, len(r))
		l = append(l, r[0])
		for _, v := range rand.Perm(len(r) - 2) {
			l = append(l, r[v+1])
		}
		l = append(l, r[len(r)-1])
		res += " " + string(l)
	}
	return res[1:]
}
