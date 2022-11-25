package main

import (
	"fmt"
	"sync"
)

/*
	Program to calculate squares of numbers using pipeline
	main thread -> value channel -> square channel -> main thread
*/

func main() {
	// array of values to calculate to find squares of using pipeline
	array := []int{2, 4, 6, 8, 10}

	// pipeline segments
	valueChannel := make(chan int)
	squareChannel := make(chan int)

	// wait group for goroutine synchronization
	wg := sync.WaitGroup{}
	wg.Add(1)

	// run goroutine for value pipeline segment
	go func() {
		for value := range valueChannel {
			squareChannel <- value * value
		}
		close(squareChannel)
	}()

	// run goroutine for square pipeline segment
	go func() {
		defer wg.Done()
		for square := range squareChannel {
			fmt.Println(square)
		}
	}()

	// iterate over array and write values to value channel
	for _, value := range array {
		valueChannel <- value
	}

	// close value channel
	close(valueChannel)

	// wait for pipleine to finish
	wg.Wait()
}
