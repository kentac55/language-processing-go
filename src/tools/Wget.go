package tools

import (
	"io/ioutil"
	"net/http"
	"os"
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
