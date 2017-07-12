package exercise

import (
	"bufio"
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
