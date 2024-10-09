package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// これでGETリクエストを送信し、レスポンスを取得することができる
	resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println("########################################################################################################")

	base, err := url.Parse("http://example.com/")
	if err != nil {
		fmt.Println(err)
		return
	}
	reference, err := base.Parse("/test?a=1&b=2")
	if err != nil {
		fmt.Println(err)
		return
	}
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(base)
	fmt.Println(endpoint)
	fmt.Println("########################################################################################################")

	// ヘッダーを追加したり
	// req, err := http.NewRequest("GET", endpoint, nil) GETリクエスト
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte("test"))) // POSTリクエスト
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	// クエリパラメータを追加したり
	q := req.URL.Query()
	q.Add("c", "3&%")
	fmt.Println(q)
	fmt.Println(q.Encode())
	req.URL.RawQuery = q.Encode()
	fmt.Println("########################################################################################################")

	var client *http.Client = &http.Client{}
	resp2, _ := client.Do(req)
	body2, _ := ioutil.ReadAll(resp2.Body)
	fmt.Println(string(body2))
	fmt.Println("########################################################################################################")
}
