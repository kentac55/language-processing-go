package exercise

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

func TestP10(t *testing.T) {
	const path = "/tmp/hightemp.txt"
	const url = "http://www.cl.ecei.tohoku.ac.jp/nlp100/data/hightemp.txt"
	rs, err := http.Get(url)
	if err != nil {
		fmt.Errorf("cannot connet target")
	}
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		fmt.Errorf("cannot read response body")
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Errorf("cannot open file")
	}
	defer file.Close()
	file.Write(body)
	out, err := exec.Command("wc", "-l", path).Output()
	if err != nil {
		t.Errorf("failure to exec command")
	}
	o, err := strconv.Atoi(strings.Split(string(out), " ")[0])
	if err != nil {
		t.Errorf("failure atoi")
	}
	x := P10(path)
	o = int(o)
	if x != o {
		t.Errorf("expect: %v, actually: %v", o, x)
	}
}
