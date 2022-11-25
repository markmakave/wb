package main

import (
	"fmt"
	"sync"
)

/*
	Implment a concurrent-safe counter
	print the value of the counter after working with it in concurrent context
*/

type Counter struct {
	value int
	mutex sync.Mutex
}

func (c *Counter) Inc() {
	// concurrent safe increment
	c.mutex.Lock()
	c.value++
	c.mutex.Unlock()
}

func main() {
	// create concurrenf-safe counter
	c := Counter{value: 0}

	// create wait group for goroutines synchronization
	wg := sync.WaitGroup{}

	// create 1000 goroutines
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			// decrement wait group counter when goroutine finishes
			defer wg.Done()

			// increment counter in concurrent context
			c.Inc()
		}()
	}

	// wait for all goroutines to finish
	wg.Wait()

	// print the counter
	fmt.Println(c.value)
}
