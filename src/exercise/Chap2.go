package exercise

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
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

func P12(p string, col1Path string, col2Path string) {
	col1Str := ""
	col2Str := ""
	f, _ := os.Open(p)
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.Split(sc.Text(), "\t")
		col1Str += line[0] + "\n"
		col2Str += line[1] + "\n"
	}
	col1File, _ := os.OpenFile(col1Path, os.O_CREATE|os.O_WRONLY, 0644)
	defer col1File.Close()
	col1File.Write([]byte(col1Str))
	col2File, _ := os.OpenFile(col2Path, os.O_CREATE|os.O_WRONLY, 0644)
	defer col2File.Close()
	col2File.Write([]byte(col2Str))
}
