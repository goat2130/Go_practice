package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func saturday() {
	fmt.Println("Welcome to Saturday!")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today is Satruday!!")
	case today + 1:
		fmt.Println("tomorrow is Saturday")
	case today + 2:
		fmt.Println("In two days is Saturday")
	default:
		fmt.Println("It's too far away")
	}
}

func greeting() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	case t.Hour() < 22:
		fmt.Println("Good evening!")
	}
}

func main() {
	defer fmt.Println("hello world!")

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwiin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	saturday()

	greeting()
}
