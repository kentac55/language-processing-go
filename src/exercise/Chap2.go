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

// -- P13 --

// P13 is a solution of `13. col1.txtとcol2.txtをマージ`
// 12で作ったcol1.txtとcol2.txtを結合し，元のファイルの1列目と2列目をタブ区切りで並べたテキストファイルを作成せよ．確認にはpasteコマンドを用いよ．

// -- P14 --

// P13 is a solution of `14. 先頭からN行を出力`
// 自然数Nをコマンドライン引数などの手段で受け取り，入力のうち先頭のN行だけを表示せよ．確認にはheadコマンドを用いよ．

// -- P15 --

// P15 is a solution of `15. 末尾のN行を出力`
// 自然数Nをコマンドライン引数などの手段で受け取り，入力のうち末尾のN行だけを表示せよ．確認にはtailコマンドを用いよ．

// -- P16 --

// P16 is a solution of `16. ファイルをN分割する`
// 自然数Nをコマンドライン引数などの手段で受け取り，入力のファイルを行単位でN分割せよ．同様の処理をsplitコマンドで実現せよ．

// -- P17 --

// P17 is a solution of `17. １列目の文字列の異なり`
// 1列目の文字列の種類（異なる文字列の集合）を求めよ．確認にはsort, uniqコマンドを用いよ．

// -- P18 --

// P18 is a solution of `18. 各行を3コラム目の数値の降順にソート`
// 各行を3コラム目の数値の逆順で整列せよ（注意: 各行の内容は変更せずに並び替えよ）．確認にはsortコマンドを用いよ（この問題はコマンドで実行した時の結果と合わなくてもよい）．

// -- P19 --

// P19 is a solution of `19. 各行の1コラム目の文字列の出現頻度を求め，出現頻度の高い順に並べる`
// 各行の1列目の文字列の出現頻度を求め，その高い順に並べて表示せよ．確認にはcut, uniq, sortコマンドを用いよ．
