package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	// content, err := ioutil.ReadFile("ioutil.go")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(string(content))

	// if err := ioutil.WriteFile("ioutil_temp.go", content, 0666); err != nil {
	// 	log.Fatalln(err)
	// }

	// bufferにバイト配列のデータを書き込む
	r := bytes.NewBuffer([]byte("abc"))
	// bufferからバイト配列のデータを読み込む
	content, _ := ioutil.ReadAll(r)
	// バイト配列のデータを文字列に変換して出力
	fmt.Println(string(content))
}
