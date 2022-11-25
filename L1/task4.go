package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
)

/*
	Create N workers to read from channel and print values to stdout.
	Exit when user presses Ctrl+C.
*/

func main() {
	// ask user for number of workers
	var n int
	fmt.Print("Enter number of workers: ")
	fmt.Scan(&n)

	// create channel for workers
	ch := make(chan int)

	// create wait group for workers synchronization
	wg := sync.WaitGroup{}

	// run n workers
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			// decrement wait group counter when goroutine finishes
			defer wg.Done()

			// run until channel is closed
			for value := range ch {
				fmt.Println(value)
			}
		}()
	}

	// create channel for OS signals
	sig := make(chan os.Signal, 1)
	// ask os to notify us about Unix SIGINT (aka Ctrl+C)
	signal.Notify(sig, os.Interrupt)

	// If we don't catch SIGINT manually, program will exit without waiting for workers to finish
	// because of default signal handler behaviour.

	// run infinite loop writing random numbers to channel
	// and exit on SIGINT
	for {
		select {
		case <-sig:
			// catched SIGINT
			// close channel to stop workers and wait for them to finish
			close(ch)
			wg.Wait()
			fmt.Println("Exited")
			return
		case ch <- rand.Int():
			// wrote random number to channel
			// can do some other stuff here
		}
	}

}
