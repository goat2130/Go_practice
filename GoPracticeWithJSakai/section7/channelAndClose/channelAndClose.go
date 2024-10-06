package main

import "fmt"

func goroutine(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
	/*
		atal error: all goroutines are asleep - deadlock!

		goroutine 1 [chan receive]:
		main.main()
			/Users/yagisawahodaka/Desktop/GoPractice/GoPracticeWithJSakai/section7/channelAndClose/channelAndClose.go:18 +0x114
		exit status 2
	*/
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, len(s)) // メモリにチャネルを確保

	go goroutine(s, c)
	// 随時チャネルに送信して値を取り出す
	for i := range c {
		fmt.Println(i)
	}
}
