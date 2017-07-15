package exercise

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"unicode/utf8"
)

func P00(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res += string(s[len(s)-i-1])
	}
	return res
}
func P01(s string) string {
	res := ""
	for i, char := range s {
		if i%2 == 0 {
			res += string(char)
		}
	}
	return res
}
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
func P03(s string) []int {
	res := []int{}
	reg := regexp.MustCompile(`[^0-9A-Za-z ]`)
	s = reg.ReplaceAllString(s, "")
	for _, v := range strings.Split(s, " ") {
		res = append(res, len(v))
	}
	return res
}
func P04(s string) map[string]int {
	res := make(map[string]int, 20)
	for i, v := range strings.Split(s, " ") {
		i++
		if i == 1 || i == 5 || i == 6 || i == 7 || i == 8 || i == 9 || i == 15 || i == 16 || i == 19 {
			res[v[:1]] = i
		} else {
			res[v[:2]] = i
		}
	}
	return res
}

type BiGram struct {
	Word   [][]string
	Letter []string
}

func P05(s string) BiGram {
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
	return BiGram{words, letters}
}

type Set map[string]struct{}

type SetFunc interface {
	Union(that Set) Set
	Intersect(that Set) Set
	Diff(that Set) Set
	Exists(that Set) bool
	Equal(that Set) bool
}

type P06Res struct {
	X         Set
	Y         Set
	union     Set
	intersect Set
	diff      Set
}

func NewSet(s []string) Set {
	set := Set{}
	if len(s) > 0 {
		for _, i := range s {
			set[i] = struct{}{}
		}
	}
	return set
}

func (s Set) Union(that Set) Set {
	set := Set{}
	for k := range s {
		set[k] = struct{}{}
	}
	for k := range that {
		set[k] = struct{}{}
	}
	return set
}

func (s Set) Intersect(that Set) Set {
	set := Set{}
	for a := range s {
		for b := range that {
			if a == b {
				set[a] = struct{}{}
			}
		}
	}
	return set
}

func (s Set) Diff(that Set) Set {
	set := Set{}
	for a := range s {
		set[a] = struct{}{}
		for b := range that {
			if a == b {
				delete(set, b)
			}
		}
	}
	return set
}

func (s Set) Exists(key string) bool {
	for k := range s {
		if k == key {
			return true
		}
	}
	return false
}

func (s Set) equal(that Set) bool {
	if len(s) != len(that) {
		return false
	}
	for a := range s {
		if _, ok := that[a]; !ok {
			return false
		}
	}
	return true
}

func P06(s1 []string, s2 []string) P06Res {
	X := NewSet(s1)
	Y := NewSet(s2)
	union := X.Union(Y)
	intersect := X.Intersect(Y)
	diff := X.Diff(Y)
	return P06Res{X, Y, union, intersect, diff}
}
func P07(x interface{}, y interface{}, z interface{}) string {
	return fmt.Sprintf("%v時の%vは%v", x, y, z)
}
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
