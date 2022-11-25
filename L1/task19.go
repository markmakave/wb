package main

import "fmt"

/*
	Reverse string
	unicode characters are supported
	done without creating a new string aka in-place
*/

func main() {
	// ask user for a string
	var s string
	fmt.Scan(&s)

	// reverse string using two indexes commig from both sides to the middle
	// i -> | <- j
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		// swap characters using some dark magic
		// because strings are immutable in Go
		s = s[:i] + string(s[j]) + s[i+1:j] + string(s[i]) + s[j+1:]
	}

	// print the result
	fmt.Println(s)
}
