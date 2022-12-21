package main

import (
	"fmt"
	"time"
)

func main() {
	// ch := make(chan int, 10) // with 10 room(buffer)

	ch := make(chan int) // with 0 room(unbuffer)

	go func() {
		for i := 0; i < 20; i++ {
			ch <- i
			fmt.Printf("thread 2 : %d\n", i)
		}

		close(ch)
	}()

	// disable this to see interasting result
	time.Sleep(2 * time.Second)

	for i := range ch {
		time.Sleep(time.Second)
		fmt.Printf("\nthread 1 : %d\n", i)
	}
}
