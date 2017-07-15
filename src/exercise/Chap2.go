package exercise

import (
	"bufio"
	"io/ioutil"
	"os"
)

func P10(p string) int {
	c := 0
	file, _ := os.Open(p)
	defer file.Close()
	sc := bufio.NewScanner(file)
	for i := &c; sc.Scan(); *i++ {
	}
	return c
}
func P11(p string) string {
	r := []byte{}
	f, _ := os.Open(p)
	defer f.Close()
	b, _ := ioutil.ReadAll(f)
	for _, v := range b {
		if v == '\t' {
			v = ' '
		}
		r = append(r, v)
	}
	return string(r)
}
