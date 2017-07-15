package exercise

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"unicode/utf8"
)

// -- P00 --

// P00 is a solution of `00. 文字列の逆順`
// 文字列"stressed"の文字を逆に（末尾から先頭に向かって）並べた文字列を得よ．
func P00(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res += string(s[len(s)-i-1])
	}
	return res
}

// -- P01 --

// P01 is a solution of `01. 「パタトクカシーー」`
//「パタトクカシーー」という文字列の1,3,5,7文字目を取り出して連結した文字列を得よ．
func P01(s string) string {
	res := ""
	for i, char := range s {
		if i%2 == 0 {
			res += string(char)
		}
	}
	return res
}

// -- P02 --

// P02 is a solution of `02. 「パトカー」＋「タクシー」＝「パタトクカシーー」`
// 「パトカー」＋「タクシー」の文字を先頭から交互に連結して文字列「パタトクカシーー」を得よ．
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

// -- P03 --

// P03 is a solution of `03. 円周率`
// "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."という文を単語に分解し，各単語の（アルファベットの）文字数を先頭から出現順に並べたリストを作成せよ．
func P03(s string) []int {
	res := []int{}
	reg := regexp.MustCompile(`[^0-9A-Za-z ]`)
	s = reg.ReplaceAllString(s, "")
	for _, v := range strings.Split(s, " ") {
		res = append(res, len(v))
	}
	return res
}

// -- P04 --

// P04 is a solution of `04. 元素記号`
// "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."という文を単語に分解し，1, 5, 6, 7, 8, 9, 15, 16, 19番目の単語は先頭の1文字，それ以外の単語は先頭に2文字を取り出し，取り出した文字列から単語の位置（先頭から何番目の単語か）への連想配列（辞書型もしくはマップ型）を作成せよ．
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

// -- P05 --

// Bigrams is a struct that has two types of n-gram
type Bigrams struct {
	Word   [][]string
	Letter []string
}

// P05 is a solution of `05. n-gram`
// 与えられたシーケンス（文字列やリストなど）からn-gramを作る関数を作成せよ．この関数を用い，"I am an NLPer"という文から単語bi-gram，文字bi-gramを得よ．
func P05(s string) Bigrams {
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
	return Bigrams{words, letters}
}

// -- P06 --

// Set is a type of Set
type Set map[string]struct{}

// SetFunc is a interface that belong to Set
type SetFunc interface {
	Union(that Set) Set
	Intersect(that Set) Set
	Diff(that Set) Set
	Exists(that Set) bool
	Equal(that Set) bool
}

// P06Res is a struct that is used by R06 only
type P06Res struct {
	X         Set
	Y         Set
	union     Set
	intersect Set
	diff      Set
}

// NewSet is a construct method of Set
func NewSet(s []string) Set {
	set := Set{}
	if len(s) > 0 {
		for _, i := range s {
			set[i] = struct{}{}
		}
	}
	return set
}

// Union is a method that returns an united Set
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

// Intersect is a method that returns an intersected Set
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

// Diff is a method that returns a diffed Set
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

// Exists is a method that returns the specific key exists in this Set or not
func (s Set) Exists(key string) bool {
	for k := range s {
		if k == key {
			return true
		}
	}
	return false
}

// equal is a method that returns the specific Set equals this Set or not
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

// -- P06 --

// P06 is a solution of `06. 集合`
// "paraparaparadise"と"paragraph"に含まれる文字bi-gramの集合を，それぞれ, XとYとして求め，XとYの和集合，積集合，差集合を求めよ．さらに，'se'というbi-gramがXおよびYに含まれるかどうかを調べよ．
func P06(s1 []string, s2 []string) P06Res {
	X := NewSet(s1)
	Y := NewSet(s2)
	union := X.Union(Y)
	intersect := X.Intersect(Y)
	diff := X.Diff(Y)
	return P06Res{X, Y, union, intersect, diff}
}

// -- P07 --

// P07 is a solution of `07. テンプレートによる文生成`
// 引数x, y, zを受け取り「x時のyはz」という文字列を返す関数を実装せよ．さらに，x=12, y="気温", z=22.4として，実行結果を確認せよ．
func P07(x interface{}, y interface{}, z interface{}) string {
	return fmt.Sprintf("%v時の%vは%v", x, y, z)
}

// -- P08 --

// P08 is a solution of `08. 暗号文`
// 与えられた文字列の各文字を，以下の仕様で変換する関数cipherを実装せよ．
//
// 英小文字ならば(219 - 文字コード)の文字に置換
// その他の文字はそのまま出力
//
// この関数を用い，英語のメッセージを暗号化・復号化せよ．
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

// -- P09 --

// P09 is a solution of `09. Typoglycemia`
// スペースで区切られた単語列に対して，各単語の先頭と末尾の文字は残し，それ以外の文字の順序をランダムに並び替えるプログラムを作成せよ．ただし，長さが４以下の単語は並び替えないこととする．適当な英語の文（例えば"I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."）を与え，その実行結果を確認せよ．
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
