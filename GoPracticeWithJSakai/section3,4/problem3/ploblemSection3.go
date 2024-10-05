package main

import "fmt"

func main() {
	P1()
	P2()
	P3()
}

func P1() {
	f := 1.11
	i := int(f)
	fmt.Println(i)
}

func P2() {
	fmt.Println("5,6")
}

func P3() {
	m := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"Messi": 30,
	}

	fmt.Printf("%T %v\n", m, m)
}
