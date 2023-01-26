package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int)

	for id := 0; id < 2; id++ {
		// each processor prints the data besides its id
		go func(id int) {
			for {
				i, ok := <-ch
				if !ok {
					fmt.Printf("Closed the processor %d\n", id)
					return
				}
				fmt.Printf("id: %d data: %d \n", id, i)
			}
		}(id)
	}

	// produce data every 1 second into the channel
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("Input")
			ch <- rand.Intn(10)
		}
	}()

	// wait for termination
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	close(ch)

	//how we can wait for processors to quit?
}
