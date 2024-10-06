package main

import "fmt"

func goroutine(s []string, c chan string) {
	defer close(c)
	sum := ""
	for _, v := range s {
		sum += v
		c <- sum
	}
}

func main() {
	words := []string{"teste1!", "teste2!", "teste3!", "teste4!"}
	c := make(chan string, len(words))
	go goroutine(words, c)
	for w := range c {
		fmt.Println(w)
	}
}
