package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *Counter) Inc(Key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[Key]++
}

func (c *Counter) Value(Key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[Key]
}

func main() {
	// c := make(map[string]int)
	c := Counter{v: make(map[string]int)}
	go func() {
		for i := 0; i < 10; i++ {
			// c["key"]++
			c.Inc("key")
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			// c["key"]++
			c.Inc("key")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(c, c.Value("key"))

}
