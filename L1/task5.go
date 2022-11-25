package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	Send random numbers to channel from writing goroutine.
	Receive numbers from channel in reading goroutine.
	Exit after N seconds.
*/

func main() {
	// ask user for number of seconds to run
	var n int
	fmt.Print("Enter a number of seconds to wait: ")
	fmt.Scan(&n)

	// create channels for workers
	quit := make(chan struct{})
	ch := make(chan int)

	// create wait group for workers synchronization
	var wg sync.WaitGroup
	wg.Add(1)

	// writing goroutine
	go func() {
	loop:
		for {
			// exit if quit channel is closed
			select {
			case <-quit:
				// main thread closed quit channel
				// close channel to stop reading goroutine
				close(ch)
				break loop
			case ch <- rand.Int():
				// write random number to channel
			}
		}
	}()

	// reading goroutine
	go func() {
		// decrement wait group counter when goroutine finishes
		defer wg.Done()

		// run until channel is closed
		for val := range ch {
			fmt.Println(val)
		}
	}()

	// wait N seconds
	time.Sleep(time.Duration(n) * time.Second)

	// close quit channel to stop writing goroutine
	close(quit)

	// wait for goroutines to finish
	wg.Wait()
}
