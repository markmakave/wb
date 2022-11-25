package main

import (
	"fmt"
	"time"
)

// custom Sleep function using time.After and channel blockng logic
func Sleep(n int) {
	<-time.After(time.Duration(n) * time.Second)
}

// usage example
func main() {
	// ask user for a number of seconds to sleep
	var n int
	fmt.Print("Enter number of seconds to sleep: ")
	fmt.Scan(&n)

	// sleep for n seconds
	Sleep(n)
}
