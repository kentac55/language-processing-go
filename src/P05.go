package exercise

import "strings"

type biGram struct {
	Word   [][]string
	Letter []string
}

func P05(s string) biGram {
	slice := strings.Split(s, " ")
	words := [][]string{}
	letters := []string{}
	for i, v := range slice {
		if i >= len(slice)-1 {
			break
		}
		words = append(words, []string{v, slice[i+1]})
	}
	for _, v := range strings.Split(s, " ") {
		if len(v) < 2 {
			continue
		}
		for i := range v {
			if i >= len(v)-1 {
				break
			}
			letters = append(letters, v[i:i+2])
		}
	}
	return biGram{words, letters}
}
