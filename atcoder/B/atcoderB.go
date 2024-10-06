package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	S := scanner.Text()

	positions := make(map[byte]int)
	for i, c := range S {
		positions[byte(c)] = i + 1 // 1-indexed
	}

	totalDistance := 0
	currentPos := positions['A']
	for c := 'B'; c <= 'Z'; c++ {
		nextPos := positions[byte(c)]
		totalDistance += abs(nextPos - currentPos)
		currentPos = nextPos
	}

	fmt.Println(totalDistance)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
