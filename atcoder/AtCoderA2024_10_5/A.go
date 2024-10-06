package main

import (
	"fmt"
)

func main() {
	var s string

	fmt.Scanf("%s", &s)

	if len(s) >= 3 && s[len(s)-3:] == "san" {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
