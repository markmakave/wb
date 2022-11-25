package main

import "fmt"

/*
	Calculate squares of array elements using goroutines (concurrency).

	- Synchronization is done using channel.
	- Results are printed in random order from main goroutine.
*/

func main() {
	// array of values to calculate squares of
	array := []int{2, 4, 6, 8, 10}

	// channel for goroutines synchronization and results
	ch := make(chan int)

	// iterate over array
	for _, v := range array {
		// run goroutine for each element
		go func(v int) {
			// calculate square of element
			var square int = v * v

			// put result to channel
			ch <- square
		}(v)
	}

	// read calculated value from channel len(array) times for synchronization
	for range array {
		// print value from channel
		fmt.Println(<-ch)
	}
}
