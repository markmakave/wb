package main

import (
	"fmt"
	"strings"
)

/*
	Reverse words in a string
*/

func main() {
	var s string = "snow dog sun"

	// split string into slice of words
	words := strings.Split(s, " ")

	// reverse the words using two indexes comming from both sides to the middle like in task19.go
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		// magic here is not needed because slices are mutable in Go
		words[i], words[j] = words[j], words[i]
	}

	// join the words and separate them with spaces
	s = strings.Join(words, " ")

	// print the result
	fmt.Println(s)
}
