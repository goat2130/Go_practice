package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	sc := bufio.NewScanner(os.Stdin)

	var us []float64
	var vs []float64

	for i := 0; i < n; i++ {
		sc.Scan()
		inputs := strings.Split(sc.Text(), " ")
		u, _ := strconv.ParseFloat(inputs[0], 64)
		v, _ := strconv.ParseFloat(inputs[1], 64)
		us = append(us, u)
		vs = append(vs, v)
	}

	var sum float64
	sum = math.Sqrt(math.Pow(0-us[0], 2) + math.Pow(0-vs[0], 2))
	for i := 0; i < n-1; i++ {
		sum += math.Sqrt(math.Pow(us[i]-us[i+1], 2) + math.Pow(vs[i]-vs[i+1], 2))
	}
	sum += math.Sqrt(math.Pow(us[n-1]-0, 2) + math.Pow(vs[n-1]-0, 2))

	fmt.Println(sum)

}
