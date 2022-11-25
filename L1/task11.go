package main

import "fmt"

/*
	Program to merge two sets of values

	Sets can be presented as maps with empty structs as values or as slices with unique values
*/

func main() {

	// Sets as maps with empty structs as values version
	{
		// given sets of values
		set1 := map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {}, 10: {}}
		set2 := map[int]struct{}{6: {}, 7: {}, 8: {}, 9: {}, 10: {}, 11: {}, 12: {}, 13: {}, 14: {}, 15: {}}

		// merged set
		merged := make(map[int]struct{})

		// merge sets
		for key := range set1 {
			merged[key] = struct{}{}
		}
		for key := range set2 {
			merged[key] = struct{}{}
		}

		fmt.Println(merged)
	}

	// Sets as slices version
	{
		// given sets of values
		set1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		set2 := []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

		// merged set
		merged := make(map[int]struct{})

		// merge sets
		for _, value := range set1 {
			merged[value] = struct{}{}
		}
		for _, value := range set2 {
			merged[value] = struct{}{}
		}

		// convert merged set to slice
		mergedSlice := make([]int, 0, len(merged))
		for key := range merged {
			mergedSlice = append(mergedSlice, key)
		}

		fmt.Println(mergedSlice)
	}

}
