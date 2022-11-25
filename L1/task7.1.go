package main

import (
	"fmt"
	"sync"
)

/*
	Write to map using goroutines (concurrency).

	- Synchronization is done using wait group.
	- Map writing is done in critical section created using mutex.
*/

func main() {
	// map to write to
	var m = make(map[string]int)

	// array of keys to write to map
	var array = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	// map writing is not thread safe
	// so we need to use mutex
	var mutex sync.Mutex

	// wait group for goroutines synchronization
	wg := sync.WaitGroup{}

	// add number of goroutines to wait group
	wg.Add(len(array))

	// iterate over array
	for i, v := range array {
		// run goroutine for each element
		go func(i int, v string) {
			// decrement wait group counter when goroutine finishes
			defer wg.Done()

			// create critical section using mutex
			mutex.Lock()
			{
				// write to map in critical section
				m[v] = i
			}
			mutex.Unlock()
		}(i, v)
	}

	// wait for all goroutines to finish
	wg.Wait()

	// print result
	fmt.Println(m)
}
