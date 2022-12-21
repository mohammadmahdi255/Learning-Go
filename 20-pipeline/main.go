package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int)

	for id := 0; id < 2; id++ {
		// each processor prints the data besides its id
		go func(id int) {
			for {
				i := <-ch
				fmt.Printf("id: %d data: %d \n", id, i)
			}
		}(id)
	}

	// produce data every 1 second into the channel
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("Input")
		ch <- rand.Intn(10)
	}

}
