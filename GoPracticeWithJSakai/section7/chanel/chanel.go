package main

import "fmt"

func goroutine(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int) // 15, 15

	go goroutine(s, c)
	go goroutine2(s, c)
	// この場合、goroutineが終わるまで待つ
	x := <-c
	fmt.Println(x)
	y := <-c
	fmt.Println(y)

	ch2 := make(chan int, 2) // バッファードチャネル saize 2
	ch2 <- 100
	fmt.Println(len(ch2)) //1
	ch2 <- 200
	fmt.Println(len(ch2)) //2

	close(ch2)

	for c := range ch2 {
		fmt.Println(c)
	}

}
