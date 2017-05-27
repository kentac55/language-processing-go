package exercise

func P01(s string) string {
	res := ""
	for i, char := range s {
		if i%2 == 0 {
			res += string(char)
		}
	}
	return res
}
