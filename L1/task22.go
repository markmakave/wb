package main

import (
	"fmt"
	"math/big"
)

/*
	Program to add, subtract, multiply and divide two numbers bigger than 2^20
*/

func main() {
	// create two bigfloats
	a := big.NewFloat(0)
	b := big.NewFloat(0)

	// set the values
	a.SetString("12371072875339021027987979982208375.902465101357402524205838")
	b.SetString("46376937677434432290009712648124896.970078050417018260538")

	// result of arithmetic operations
	c := big.NewFloat(0)

	// add the two numbers
	c.Add(a, b)
	fmt.Println(c.String())

	// subtract the two numbers
	c.Sub(a, b)
	fmt.Println(c.String())

	// multiply the two numbers
	c.Mul(a, b)
	fmt.Println(c.String())

	// divide the two numbers
	c.Quo(a, b)
	fmt.Println(c.String())
}
