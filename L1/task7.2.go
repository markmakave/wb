package main

import "fmt"

/*
	Write to map using goroutines (concurrency).

	- Synchronization is done using channel.
	- Map writing is done in critical section created using channel.
*/

func main() {
	// map to write to
	var m = make(map[string]int)

	// array of keys to write to map
	var array = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	// map writing is not thread safe
	// so we need to use mutex-like channel
	var mutex = make(chan struct{}, 1)

	// channel for goroutines synchronization
	var ch = make(chan struct{})

	// iterate over array
	for i, v := range array {
		// run goroutine for each element
		go func(i int, v string) {
			// create critical section using mutex-like channel
			<-mutex
			{
				// write to map in critical section
				m[v] = i
			}
			mutex <- struct{}{}

			// tell the channel that goroutine finished
			ch <- struct{}{}
		}(i, v)
	}

	// put one value to mutex channel to unlock it
	mutex <- struct{}{}

	// wait for all goroutines to finish
	for range array {
		<-ch
	}

	// print result
	fmt.Println(m)
}
