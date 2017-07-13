package exercise

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"testing"
)

func TestP11(t *testing.T) {
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
	o, err := exec.Command("sed", "-r", "-e", "s/\t/ /g", path).Output()
	if err != nil {
		t.Errorf("failure to exec command")
	}
	if x := P11(path); x != string(o) {
		t.Errorf("expect: %v, actually: %v", o, x)
	}
}
