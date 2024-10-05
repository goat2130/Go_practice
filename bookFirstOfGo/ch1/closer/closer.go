package main

import "fmt"

func incrementGenerator() func() int {
	var x int = 0
	return func() int {
		x++
		return x
	}
}

func circleAreaGenerator(pi float64) func(float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func main() {
	increment := incrementGenerator()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	circleArea := circleAreaGenerator(3.14)
	fmt.Println(circleArea(2))
	fmt.Println(circleArea(3))
	/*
		12.56
	    28.259999999999998
	*/

	circleArea2 := circleAreaGenerator(3)
	fmt.Println(circleArea2(2))
	fmt.Println(circleArea2(3))
	/*
		12
		27
	*/
}
