package exercise

import (
	"os/exec"
	"testing"
)

func TestP11(t *testing.T) {
	const path = "../hightemp.txt"
	o, e := exec.Command("sed", "-r", "-e", "s/\t/ /g", path).Output()
	if e != nil {
		t.Errorf("failure to exec command")
	}
	if x := P11(path); x != string(o) {
		t.Errorf("expect: %v, actually: %v", o, x)
	}
}
