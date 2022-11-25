package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	a 1000 and 1 ways to kill a goroutine (3 actually)

	use a channel to signal the goroutine to exit
*/

func main() {
	// crate wait group for goroutine synchronization
	var wg sync.WaitGroup

	// create channel for telling goroutine to exit by closing it
	quit := make(chan struct{})

	// run goroutine
	wg.Add(1)
	go func() {
		// decrement wait group counter when goroutine finishes
		defer wg.Done()

		// run until quit channel is closed
	loop:
		for {
			select {
			case <-quit:
				// main thread closed quit channel
				break loop
			default:
				// do goroutine stuff
			}
		}

		fmt.Println("Quitting...")
	}()

	// do main thread stuff
	time.Sleep(2 * time.Second)

	// tell goroutine to exit
	close(quit)

	// wait for goroutine to finish
	wg.Wait()
}
