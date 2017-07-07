package exercise

import "strings"

type biGram struct {
	Word   []wordBiGram
	Letter []letterBiGram
}

type wordBiGram struct {
	x string
	y string
}

type letterBiGram struct {
	v string
}

func P05(s string) biGram {
	slice := strings.Split(s, " ")
	words := []wordBiGram{}
	letters := []letterBiGram{}
	for i, v := range slice {
		if i >= len(slice)-1 {
			break
		}
		words = append(words, wordBiGram{v, slice[i+1]})
	}
	for i := range s {
		if i >= len(s)-1 {
			break
		}
		letters = append(letters, letterBiGram{s[i : i+2]})
	}
	return biGram{words, letters}
}
