package exercise

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

// -- P10 --

// P10 is a solution of `10. 行数のカウント`
// 行数をカウントせよ．確認にはwcコマンドを用いよ．
func P10(p string) int {
	c := 0
	file, _ := os.Open(p)
	defer file.Close()
	sc := bufio.NewScanner(file)
	for i := &c; sc.Scan(); *i++ {
	}
	return c
}

// -- P11 --

// P11 is a solution of `11. タブをスペースに置換`
//タブ1文字につきスペース1文字に置換せよ．確認にはsedコマンド，trコマンド，もしくはexpandコマンドを用いよ．
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

// -- P12 --

// P12 is a solution of `12. 1列目をcol1.txtに，2列目をcol2.txtに保存`
//各行の1列目だけを抜き出したものをcol1.txtに，2列目だけを抜き出したものをcol2.txtとしてファイルに保存せよ．確認にはcutコマンドを用いよ．
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
