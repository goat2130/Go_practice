package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

// 同時にいくつのgoroutineを実行するかを制御する
var s *semaphore.Weighted = semaphore.NewWeighted(1)

func longProcess(ctx context.Context) {
	// 1つのgoroutineが実行中の場合は、他のgoroutineが走った場合キャンセルとする
	isAcquire := s.TryAcquire(1)
	if !isAcquire {
		fmt.Println("Could not get lock")
		return
	}

	// 並列処理を制限する場合は、Acquireを使う
	// if err := s.Acquire(ctx, 1); err != nil {
	// 	fmt.Println("Could not get lock")
	// 	return
	// }
	defer s.Release(1)
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func main() {
	ctx := context.TODO()
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	time.Sleep(2 * time.Second)
	go longProcess(ctx)
	time.Sleep(5 * time.Second)
}
