package language_processing_go

func P00(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res += string(s[len(s)-i-1])
	}
	return res
}
