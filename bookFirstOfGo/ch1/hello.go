package main

import "fmt"

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
	runeKey = 'a'
)

const z = 20 * 10

func main() {
	// fmt.Println("Hello, World!")
	// var x int
	// x = 10
	// fmt.Println(x)

	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	x = x + 10
	y = "world"

	fmt.Println(x)
	fmt.Println(y)
}
