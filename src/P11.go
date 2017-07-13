package exercise

import (
	"io/ioutil"
	"os"
)

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
