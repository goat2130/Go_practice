package main

import (
	"context"
	"fmt"
	"time"
)

func longProcess(ch chan string, ctx context.Context) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

func main() {
	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	// ctx := context.TODO()　何もしないコンテキスト、将来的に使うことになるかもしれないが現時点でcontextに何を入れるかわからない場合に使う
	defer cancel()
	go longProcess(ch, ctx)

CTXLOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break CTXLOOP
		case <-ch:
			fmt.Println("success")
			break CTXLOOP
		}
	}
	fmt.Println("################")
}
