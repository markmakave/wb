package main

import "fmt"

/*
	Swap two variables without using a third variable.
*/

func main() {
	var a, b int = 100, 200

	// swap using oneliner
	a, b = b, a
	fmt.Println(a, b)

	// swaps below are not overflow-safe

	// swap using add
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)

	// swap using sub
	a = a - b
	b = a + b
	a = b - a
	fmt.Println(a, b)

	// swap using mul
	a = a * b
	b = a / b
	a = a / b
	fmt.Println(a, b)

	// swap usig xor
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)
}
