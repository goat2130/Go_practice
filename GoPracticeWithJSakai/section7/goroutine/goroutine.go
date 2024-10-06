package main

import (
	"fmt"
	"sync"
)

// 軽量なスレッド、マルチスレッド、マルチプロセス、イベント駆動などとも呼ばれる。一般的には並列処理と呼ばれる。
func goroutine(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup

	// waitGroupに1つの並列処理を追加、あることを明示的に示す
	wg.Add(1)

	go goroutine("World", &wg)
	normal("Hello")

	// 並列処理のwg.done()が終わるまで待つ。
	wg.Wait()
}
