package main

import "fmt"

/*
	Program to set or clear a bit in a number
*/

func main() {
	// ask user to enter number to modify
	var value int64
	fmt.Print("Enter value to modify: ")
	fmt.Scan(&value)

	// ask user to enter index of modified bit
	var i int
	fmt.Print("Enter a bit index to modify: ")
	fmt.Scan(&i)

	// ask user to enter new bit value
	var bit bool
	fmt.Print("Enter a bit value to set: ")
	fmt.Scan(&bit)

	// modify bit
	if bit {
		value |= 1 << i
	} else {
		value &= ^(1 << i)
	}

	// print result
	fmt.Println(value)
}
