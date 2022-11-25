package main

import "fmt"

/*
	Make a set of given strings
*/

func main() {
	// given strings
	slice := []string{"cat", "cat", "dog", "cat", "tree"}

	// create map to store unique values
	set := make(map[string]struct{})

	// iterate over slice
	for _, value := range slice {
		// add value to the set
		set[value] = struct{}{}

		// if value is already in the set, nothing will happen
	}

	// print resulting set
	fmt.Println(set)
}
