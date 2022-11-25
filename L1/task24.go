package main

import (
	"fmt"
	"math"
)

/*
	Implement 2D Point struct with incapsulated fields x and y and constructor
	Implement a function that calculates distance between two points
*/

// 2D Point structure
type Point struct {
	x, y float64
}

// Point constructor
func newPoint(x, y float64) Point {
	return Point{x, y}
}

// function that calculates distance between two points using Pythagorean theorem
func distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func main() {
	// create two points
	p1 := newPoint(1.0, 2.0)
	p2 := newPoint(3.0, 4.0)

	// calculate distance between them
	d := distance(p1, p2)

	// print the result
	fmt.Println(d)
}
