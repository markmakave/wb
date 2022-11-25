package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/*
	Implement binary search using 'built-in language tools'
*/

/*
	If we talk about the built-in tools of the language
	the whole algorithm section below can be replaced with:

	index := sort.SearchInts(array, value)
	fmt.Println(index)

	it will tell us the index of the value in the array if it is present
	otherwise it will tell us the place where the value should be inserted

	but it is too boring :)
*/

func main() {

	// create array of random numbers
	array := make([]int, 10)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		array[i] = rand.Intn(100)
	}

	sort.Ints(array)
	fmt.Println(array)

	var value int
	fmt.Scan(&value)

	// binary search
	i := 0
	j := len(array) - 1
	for i <= j {
		m := (i + j) / 2
		if array[m] == value {
			// found
			fmt.Println(m)
			return
		} else if array[m] < value {
			// search right
			i = m + 1
		} else {
			// search left
			j = m - 1
		}
	}

	// not found
	fmt.Println(-1)
}
