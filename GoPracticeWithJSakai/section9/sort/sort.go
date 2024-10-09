package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	s := []string{"peach", "banana", "kiwi"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Milke", 23},
		{"Alice", 21},
		{"Bob", 25},
		{"Nancy", 20},
	}

	fmt.Println(i, s, p)

	sort.Ints(i)
	sort.Strings(s)
	// sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
	sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age })
	fmt.Println(i, s, p)
}
