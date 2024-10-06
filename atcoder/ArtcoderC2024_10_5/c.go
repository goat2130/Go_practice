package main

import (
	"fmt"
	"math"
)

func minMaxLunchGroup(n int, K []int) int {
	const INF = int(2e9)
	total := 0
	for _, k := range K {
		total += k
	}

	ans := INF

	// ビット全探索
	for i := 0; i < (1 << n); i++ {
		groupA := 0
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				groupA += K[j]
			}
		}
		groupB := total - groupA
		ans = int(math.Min(float64(ans), float64(math.Max(float64(groupA), float64(groupB)))))
	}

	return ans
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	K := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&K[i])
	}

	result := minMaxLunchGroup(n, K)
	fmt.Println(result)
}
