package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
	a 1000 and 1 ways to kill a goroutine (3 actually)

	use a atomic captured variable to tell goroutine to stop
*/

func main() {
	// crate wait group for goroutine synchronization
	var wg sync.WaitGroup

	// create variable for telling goroutine to exit
	var shouldQuit atomic.Bool
	shouldQuit.Store(false)

	// run goroutine
	wg.Add(1)
	go func() {
		// decrement wait group counter when goroutine finishes
		defer wg.Done()

		// run until shouldQuit is true
		for !shouldQuit.Load() {
			// do goroutine stuff
		}

		fmt.Println("Quitting...")
	}()

	// do main thread stuff
	time.Sleep(2 * time.Second)

	// tell goroutine to exit
	shouldQuit.Store(true)

	// wait for goroutine to finish
	wg.Wait()
}
