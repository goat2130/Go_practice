package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	K := make([]int, n)
	total := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&K[i])
		total += K[i]
	}

	dp := make([]bool, total/2+1)
	dp[0] = true

	for _, k := range K {
		for j := total / 2; j >= k; j-- {
			dp[j] = dp[j] || dp[j-k]
		}
	}

	// total/2 に最も近い部分和を探す
	half := total / 2
	for half >= 0 && !dp[half] {
		half--
	}

	// グループ A の人数の合計が half であるとき、グループ B の人数の合計は total - half
	groupA := half
	groupB := total - half

	// 同時に昼休みをとる最大人数としてあり得る最小の値
	result := int(math.Max(float64(groupA), float64(groupB)))
	fmt.Println(result)
}
