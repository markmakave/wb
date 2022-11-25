package main

import (
	"fmt"
	"sync"
)

/*
	Calculate squares of array elements using goroutines (concurrency).

	- Synchronization is done using wait group.
	- Results are printed in random order from worker goroutines.
*/

func main() {
	// given array of integers
	array := []int{2, 4, 6, 8, 10}

	// wait group for goroutines synchronization
	wg := sync.WaitGroup{}

	// add number of goroutines to wait group
	wg.Add(len(array))

	// iterate over array
	for _, v := range array {
		// run goroutine for each element
		go func(v int) {
			// decrement wait group counter when goroutine finishes
			defer wg.Done()

			// calculate square of element
			var square int = v * v

			// print result
			fmt.Println(square)
		}(v)
	}

	// wait for all goroutines to finish
	wg.Wait()
}
