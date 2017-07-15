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

func TestP12(t *testing.T) {
	const path = "/tmp/hightemp.txt"
	const url = "http://www.cl.ecei.tohoku.ac.jp/nlp100/data/hightemp.txt"
	if err := MyWget(path, url); err != nil {
		t.Errorf("Error on wget: %v", err)
	}
	const col1Path = "/tmp/col1Str.txt"
	const col2Path = "/tmp/col2Str.txt"
	P12(path, col1Path, col2Path)
	expect := ""
	actually := ""

	// COLUMN 1
	col1Output, err := exec.Command("cut", "-f", "1", path).Output()
	if err != nil {
		t.Errorf("failure to exec command")
	}
	col1File, err := os.Open(col1Path)
	if err != nil {
		t.Errorf("cannot open: %v", col1Path)
	}
	col1Byte, err := ioutil.ReadAll(col1File)
	if err != nil {
		t.Errorf("cannot read binary: %v", col1Path)
	}
	expect = string(col1Output)
	actually = string(col1Byte)
	if expect != actually {
		t.Errorf("expect: \n%v\nactually: \n%v", expect, actually)
	}

	// COLUMN 2
	col2Output, err := exec.Command("cut", "-f", "2", path).Output()
	if err != nil {
		t.Errorf("failure to exec command")
	}
	col2File, err := os.Open(col2Path)
	if err != nil {
		t.Errorf("cannot open: %v", col2Path)
	}
	col2Byte, err := ioutil.ReadAll(col2File)
	if err != nil {
		t.Errorf("cannot read binary: %v", col2Path)
	}
	expect = string(col2Output)
	actually = string(col2Byte)
	if expect != actually {
		t.Errorf("expect: \n%v\nactually: \n%v", expect, actually)
	}
}

func TestP13(t *testing.T) {
	const col1Path = "/tmp/col1Str.txt"
	const col2Path = "/tmp/col2Str.txt"
	const concatenatePath = "/tmp/concatenate.txt"
	P13(concatenatePath, col1Path, col2Path)
	o, err := exec.Command("paste", col1Path, col2Path).Output()
	if err != nil {
		t.Errorf("failure to exec command")
	}
	concatenateFile, err := os.Open(concatenatePath)
	if err != nil {
		t. Errorf("cannot open: %v", concatenatePath)
	}
	concatenateByte, err := ioutil.ReadAll(concatenateFile)
	if err != nil {
		t.Errorf("cannot read binary: %v", concatenatePath)
	}
	expect := string(o)
	actually := string(concatenateByte)
	if expect != actually {
		t.Errorf("expect: \n%v\nactually: \n%v", expect, actually)
	}
}
