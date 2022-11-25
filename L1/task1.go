package main

import "fmt"

/*
	- Create a struct Human with some fields.
	- Create a struct Action embedded with Human.
*/

// Base Human struct
type Human struct {
	// Some Human fields
	Age  int
	Name string
}

// Some Human methods
func (h *Human) Birthsday() {
	h.Age++
}
func (h *Human) Rename(name string) {
	h.Name = name
}

// Action struct with Human embedded methods and fields
type Action struct {
	Human
}

// Try using Human methods and fields on Action
func main() {

	a := Action{Human{Age: 10, Name: "John"}}

	a.Birthsday()
	a.Rename("Jack")

	fmt.Println(a)
}
