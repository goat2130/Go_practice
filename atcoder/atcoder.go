package main

import (
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner((os.Stdin))

	var strings []string
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	var result int
	for i := 1; i <= 12; i++ {
		if len(strings[i-1]) == i {
			result++
		}
	}
}
