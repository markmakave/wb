package main

import (
	"fmt"
	"math"
)

/*
	Split values of float array in groups with step of 10.0

	Two solutions are presented:
		- Splitting values only by number of tens in it
		- Splitting values by number of tens in it and by sign (for handling numbers in (-10.0, 10.0) range)
*/

func main() {

	// Splitting values only by number of tens in it
	// warning: values int (-10.0, 10.0) are grouped in one group

	// given tempreature values
	tempreatureValues := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// resulting split map
	split := make(map[int][]float64)

	// loop over tempreature values
	for _, value := range tempreatureValues {
		// find number of tens in tempreature value
		key := int(value) / 10 * 10

		// insert tempreature value into split map by key
		split[key] = append(split[key], value)
	}

	fmt.Println(split)

	// Splitting values by number of tens in it and by sign
	// (-10.0, 10.0) problem is solved

	tempreatureValues = append(tempreatureValues, 9.3, -6.5)

	type Key struct {
		value    uint
		negative bool
	}

	smartSplit := make(map[Key][]float64)
	for _, value := range tempreatureValues {
		key := Key{value: uint(math.Abs(value/10)) * 10, negative: value < 0}

		smartSplit[key] = append(smartSplit[key], value)
	}

	fmt.Println(smartSplit)
}
