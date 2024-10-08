package main

import (
	"fmt"
	"time"
)

func failedUpdate(px *int) {
	x2 := 20
	px = &x2
}

func update(px *int) {
	*px = 20
}

type Counter struct {
	total      int
	lastUpdate time.Time
}

// ポインタレシーバー
func (c *Counter) Increment() {
	c.total++
	c.lastUpdate = time.Now()
}

// 値レシーバー
func (c Counter) String() string {
	return fmt.Sprintf("{total: %d, lastUpdate: %v}", c.total, c.lastUpdate)
}

// ポインタレシーバーに値をコピーして渡す
func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println(c.String())
}

func doUpdate(c *Counter) {
	c.Increment()
	fmt.Println(c.String())
}

func main() {
	x := 10
	failedUpdate(&x)
	fmt.Println(x)

	update(&x)
	fmt.Println(x)

	var counter Counter
	fmt.Println(counter.String())

	counter.Increment()
	fmt.Println(counter.String())

	doUpdateWrong(counter)

	doUpdate(&counter)
}
