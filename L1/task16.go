package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Implement quick sort using built-in language features.
*/

/*
	The whole section below can be replaced with:
	sort.Ints(slice)

	but it's too boring :)
*/

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	pivot := arr[left]
	i := left
	j := right
	for i < j {
		for i < j && arr[j] >= pivot {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}
		for i < j && arr[i] <= pivot {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = pivot
	quickSort(arr, left, i-1)
	quickSort(arr, i+1, right)
}

func main() {
	// create slice of random numbers
	slice := make([]int, 10)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		slice[i] = rand.Intn(1000)
	}

	// print slice before sorting
	fmt.Println(slice)

	// sort slice
	quickSort(slice, 0, len(slice)-1)

	// print slice after sorting
	fmt.Println(slice)
}
