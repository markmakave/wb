package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
	Fix the potential problems in the code below.


	var justString string
	func someFunc() {
		v := createHugeString(1 << 10)
		justString = v[:100]
	}

	func main() {
		someFunc()
	}
*/

// The problem is that cut off the string to only 100 characters,
// but the while 1024 characters would be stored in memory even
// after the function returns.

// So we need to manually copy first 100 runes by creating rune slice
// and then convert it back to string.

func createHugeString(size int) string {
	if size == 0 {
		return ""
	}
	return "a" + createHugeString(size-1)
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]

	// Pointers will be the same because we are using the same string
	// and this chunk of memory will remain allocated instead of being cut off.
	fmt.Println("[ WRONG ] Huge string header: ", (*reflect.StringHeader)(unsafe.Pointer(&v)))
	fmt.Println("[ WRONG ] Just string header: ", (*reflect.StringHeader)(unsafe.Pointer(&justString)))

	// Copy string properly
	justString = string([]rune(v)[:100])

	// Check if the string is copied correctly
	// Pointers should be different
	fmt.Println("[ CORRECT ] Huge string header: ", (*reflect.StringHeader)(unsafe.Pointer(&v)))
	fmt.Println("[ CORRECT ] Just string header: ", (*reflect.StringHeader)(unsafe.Pointer(&justString)))
}

func main() {
	someFunc()
}
