package language_processing_go

func P01(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res += string(s[len(s)-i-1])
	}
	return res
}
