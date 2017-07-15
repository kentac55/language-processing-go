package exercise

import (
	"github.com/kentac55/language-processing-go/src/tools"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

func TestP10(t *testing.T) {
	const path = "/tmp/hightemp.txt"
	const url = "http://www.cl.ecei.tohoku.ac.jp/nlp100/data/hightemp.txt"
	if err := tools.MyWget(path, url); err != nil {
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
	if err := tools.MyWget(path, url); err != nil {
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
