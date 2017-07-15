package exercise

import (
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

func MyWget(path string, url string) error {
	if _, err := os.Stat(path); err == nil {
		return nil
	}
	rs, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(body)
	return nil
}

func TestP10(t *testing.T) {
	const path = "/tmp/hightemp.txt"
	const url = "http://www.cl.ecei.tohoku.ac.jp/nlp100/data/hightemp.txt"
	if err := MyWget(path, url); err != nil {
		t.Errorf("Error on wget: %v", err)
	}
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

func TestP11(t *testing.T) {
	const path = "/tmp/hightemp.txt"
	const url = "http://www.cl.ecei.tohoku.ac.jp/nlp100/data/hightemp.txt"
	if err := MyWget(path, url); err != nil {
		t.Errorf("Error on wget: %v", err)
	}
	o, err := exec.Command("sed", "-r", "-e", "s/\t/ /g", path).Output()
	if err != nil {
		t.Errorf("failure to exec command")
	}
	if x := P11(path); x != string(o) {
		t.Errorf("expect: %v, actually: %v", o, x)
	}
}
