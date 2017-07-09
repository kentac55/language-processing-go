package exercise

import "strings"

type biGram struct {
	Word   [][]string
	Letter []string
}

func P05(s string) *biGram {
	slice := strings.Split(s, " ")
	words := [][]string{}
	letters := []string{}
	for i, v := range slice {
		if i >= len(slice)-1 {
			break
		}
		words = append(words, []string{v, slice[i+1]})
	}
	for i := range s {
		if i >= len(s)-1 {
			break
		}
		letters = append(letters, s[i:i+2])
	}
	return &biGram{words, letters}
}
