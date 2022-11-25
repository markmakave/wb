package main

import "fmt"

/*
	Calculate sum of squares of array elements using goroutines (concurrency).

	- Synchronization is done using channel.
	- Safe arithmetic summation is done in main goroutine.
*/

func main() {
	// array of values to calculate sum of squares of
	array := []int{2, 4, 6, 8, 10}

	// create channel
	ch := make(chan int)

	// iterate over array
	for _, v := range array {
		// run goroutine for each element
		go func(v int) {
			// calculate square of element in concurrent context
			square := v * v

			// write square of element to channel and lock until it is read
			ch <- square
		}(v)
	}

	// result sum of squares
	var sum int = 0

	// iterate over array
	for range array {
		// read value from channel and add it to sum
		sum += <-ch
	}

	// print result
	fmt.Println(sum)
}
