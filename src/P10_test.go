package exercise

import (
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

func TestP10(t *testing.T) {
	const path = "../hightemp.txt"
	out, err1 := exec.Command("wc", "-l", path).Output()
	if err1 != nil {
		t.Errorf("failure to exec command")
	}
	o, err2 := strconv.Atoi(strings.Split(string(out), " ")[0])
	if err2 != nil {
		t.Errorf("failure atoi")
	}
	x := P10(path)
	o = int(o)
	if x != o {
		t.Errorf("expect: %v, actually: %v", o, x)
	}
}
