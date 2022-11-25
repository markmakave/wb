package main

import (
	"fmt"
	"strings"
)

/*
	Implement a function that checks if a string contains only unique characters
	unicode characters are allowed
	function must be case insensitive
*/

func main() {
	// create set of runes aka int32 (characters including unicode)
	charMap := make(map[rune]struct{})

	// ask user for a string
	var s string
	fmt.Scan(&s)

	// convert string to lowercase to make it case insensitive
	strings.ToLower(s)

	// iterate over string
	for _, char := range s {
		// if character is already in the set, return false
		// add character to the set otherwise
		if _, ok := charMap[char]; !ok {
			// add character to the set
			charMap[char] = struct{}{}
		} else {
			// fount duplicate character
			fmt.Println("false")
			return
		}
	}

	// no duplicate characters found
	fmt.Println("true")
}
