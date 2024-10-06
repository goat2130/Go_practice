package main

import "fmt"

func main() {
	var s string
	var t string
	fmt.Scanf("%s", &s)
	fmt.Scanf("%s", &t)

	if s == t {
		fmt.Println(0)
		return
	}

	minLen := len(s)
	if len(t) < minLen {
		minLen = len(t)
	}

	for i := 0; i < minLen; i++ {
		if s[i] != t[i] {
			fmt.Println(i + 1)
			return
		}
	}

	// ここまで来た場合、s と t の一方が他方の接頭辞である
	fmt.Println(minLen + 1)
}
