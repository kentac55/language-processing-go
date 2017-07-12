package exercise

func P08(s string) string {
	res := []rune{}
	for _, r := range s {
		if r >= 97 && r <= 122 {
			r = rune(219 - r)
		} else {
			r = rune(r)
		}
		res = append(res, r)
	}
	return string(res)
}
