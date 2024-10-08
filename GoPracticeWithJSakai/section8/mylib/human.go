package mylib

import "fmt"

var Public string = "Public"
var private string = "private"

// 大文字のキャピタルでないと、他のパッケージから呼び出すことができない
type Person struct {
	Name string
	Age  int
}

func Say() {
	fmt.Println("Hello")
}
