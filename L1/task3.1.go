package main

import (
	"fmt"
	"sync"
)

/*
	Calculate sum of squares of array elements using goroutines (concurrency).

	- Synchronization is done using wait group.
	- Safe arithmetic summation is done using mutex.
*/

func main() {
	// array of values to calculate sum of squares of
	array := []int{2, 4, 6, 8, 10}

	// result sum of squares
	var sum int = 0

	// arithmetic sum is not thread safe
	// so we need to use mutex
	var mutex sync.Mutex

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

			// calculate square of element (still in concurrent context)
			var square int = v * v

			// create critical section
			mutex.Lock()
			{
				// add square to sum in critical section
				sum += square
			}
			mutex.Unlock()
		}(v)
	}

	// wait for all goroutines to finish
	wg.Wait()

	// print result
	fmt.Println(sum)
}
