package main

import "fmt"

/*
	Remove element from slice by index
*/

func main() {
	// given slice of integers
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// ask user for an index
	var n int
	fmt.Print("Enter index of element to delete: ")
	fmt.Scan(&n)

	// remove element from slice
	slice = append(slice[:n], slice[n+1:]...)

	// print the result
	fmt.Println(slice)
}
