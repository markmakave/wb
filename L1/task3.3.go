package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	Calculate sum of squares of array elements using goroutines (concurrency).

	- Synchronization is done using wait group.
	- Safe arithmetic summation is done using atomic operations.
*/

func main() {
	// array of values to calculate sum of squares of
	array := []int64{2, 4, 6, 8, 10}

	// result sum of squares
	var sum int64 = 0

	// wait group for goroutines synchronization
	wg := sync.WaitGroup{}

	// add number of goroutines to wait group
	wg.Add(len(array))

	// iterate over array
	for _, v := range array {
		// run goroutine for each element
		go func(v int64) {
			// decrement wait group counter when goroutine finishes
			defer wg.Done()

			// calculate square of element in concurrent context
			var square int64 = v * v

			// atomic add square to sum
			atomic.AddInt64(&sum, square)
		}(v)
	}

	// wait for all goroutines to finish
	wg.Wait()

	// print result
	fmt.Println(sum)
}
