package main

import (
	"fmt"
	"math"
	"runtime"
	"strings"
	"time"
)

type Vertex struct {
	X int
	y int
}

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

func pointer() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

func primes(primes []int) []int {
	result := primes[1:4]
	return result
}

func Slice() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "xxxx"
	fmt.Println(a, b)
	fmt.Println(names)
}

func sliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 14}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
		s string
	}{
		{2, true, "hello"},
		{3, false, "world"},
		{5, true, "Go"},
	}
	fmt.Println(s)
}

func printSliceAndCap() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("len = %d cap = %d %v\n", len(s), cap(s), s)

	s = s[:0]
	fmt.Printf("len = %d cap = %d %v\n", len(s), cap(s), s)

	s = s[:4]
	fmt.Printf("len = %d cap = %d %v\n", len(s), cap(s), s)

	s = s[2:]
	fmt.Printf("len = %d cap = %d %v\n", len(s), cap(s), s)
}

func printNil() {
	if n := []int{}; len(n) == 0 {
		fmt.Println(n, len(n), cap(n))
		fmt.Println("nil!!!")
	}
}

func creatingSliceWithMake() {
	a := make([]int, 5)
	fmt.Println(a, len(a), cap(a))

	c := make([]int, 0, 5)
	fmt.Println(c, len(c), cap(c))

	d := c[:2]
	fmt.Println(d, len(d), cap(d))

	e := d[2:5]
	fmt.Println(e, len(e), cap(e))
}

func SlicesOfSlice() {
	noard := [][]string{
		{"Go", "is", "awesome"},
		{"Go", "is", "awesome"},
		{"Go", "is", "awesome"},
	}

	fmt.Println(noard)

	noard[0][0] = "x"
	noard[0][1] = "o"
	noard[0][2] = "x"
	noard[1][0] = "o"
	noard[1][1] = "x"
	noard[1][2] = "o"
	noard[2][0] = "x"
	noard[2][1] = "o"
	noard[2][2] = "o"

	for i := 0; i < len(noard); i++ {
		fmt.Printf("%s\n", strings.Join(noard[i], " "))
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

	pointer()

	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	var (
		v1 = Vertex{1, 2} // has type Vertex
		v2 = Vertex{X: 1}
		v3 = Vertex{}
		p  = &Vertex{1, 2}
	)

	fmt.Println(v1, p, v2, v3)

	primeList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := primes(primeList)
	fmt.Println(result)

	Slice()

	sliceLiterals()

	printSliceAndCap()

	printNil()

	creatingSliceWithMake()

	SlicesOfSlice()

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %v\n", i, v)
	}

	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i)
	}

	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}
}
