package main

import (
	"fmt"
)

func main() {
	p1()
	p2()
}

func p1() {
	l := []int{100, 300, 11, 23, 2, 4, 6, 4}

	var min int
	for i, v := range l {
		if i == 0 {
			min = v
			continue
		}
		if v < min {
			min = v
		}
	}
	fmt.Println(min)
}

func p2() {
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"graph":  150,
		"orange": 150,
		"papaya": 500,
		"kiwi":   90,
	}

	result := 0
	for _, v := range m {
		result += v
	}
	fmt.Println(result)
}
