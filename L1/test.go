package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func someAction(v []int8, b int8) {
	v[0] = 100
	vn := append(v, b)

	fmt.Println(v)
	fmt.Println(vn)

	// print the slice pointer
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&v)))
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&vn)))

}

func main() {
	var a = make([]int8, 5, 10)
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5
	someAction(a, 6)
	fmt.Println(a)
}
