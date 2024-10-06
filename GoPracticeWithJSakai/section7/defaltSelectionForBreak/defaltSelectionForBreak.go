package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

OuterLoop: // This is a label for the outer loop.
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!!")
			break OuterLoop // This will exit the outer loop.
			// return
		default:
			fmt.Println("	.")
			time.Sleep(50 * time.Millisecond)
		}
	}

	fmt.Println("#################")
}
