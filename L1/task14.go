package main

import "fmt"

// function created to detect type of interface{} variable
func detect(value interface{}) {
	// try casting it to int
	if _, ok := value.(int); ok {
		fmt.Println("int")
		return
	}

	// try casting it to string
	if _, ok := value.(string); ok {
		fmt.Println("string")
		return
	}

	// try casting it to bool
	if _, ok := value.(bool); ok {
		fmt.Println("bool")
		return
	}

	// Cast it into the fire. Destroy it!
	fmt.Println("No.")
}

func main() {
	// Test function with different types
	detect(1)
	detect("1")
	detect(true)
	detect(1.0)
}
