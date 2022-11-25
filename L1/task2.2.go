package main

import "fmt"

/*
	Calculate squares of array elements using goroutines (concurrency).

	- Synchronization is done using channel.
	- Results are printed in random order from worker goroutines.
*/

func main() {
	// array of values to write to channel
	array := []int{2, 4, 6, 8, 10}

	// create channel
	ch := make(chan struct{})

	// iterate over array
	for _, v := range array {
		// run goroutine for each element
		go func(v int) {
			// calculate square of element
			var square int = v * v

			// print result
			fmt.Println(square)

			// tell the channel that goroutine finished
			ch <- struct{}{}
		}(v)
	}

	// read some value from channel len(array) times for synchronization
	for range array {
		<-ch
	}
}
