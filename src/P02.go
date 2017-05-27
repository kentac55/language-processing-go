package exercise

import (
	"errors"
	"unicode/utf8"
)

func P02(s1 string, s2 string) (string, error) {
	res := ""
	if utf8.RuneCountInString(s1) != utf8.RuneCountInString(s2) {
		return "", errors.New("not match str length")
	}
	for i, charA := range s1 {
		for j, charB := range s2 {
			if i == j {
				res += string(charA) + string(charB)
			}
		}
	}
	return res, nil
}
