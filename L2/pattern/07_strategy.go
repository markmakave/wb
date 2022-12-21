package main

import (
	"fmt"
)

type Strategy interface {
	execute()
}

type QuickSort struct {
}

func (q *QuickSort) execute() {
	fmt.Println("Quick sort")
}

type InsertionSort struct {
}

func (i *InsertionSort) execute() {
	fmt.Println("Insertion sort")
}

type Sorter struct {
	strategy Strategy
}

func (s *Sorter) setStrategy(strategy Strategy) {
	s.strategy = strategy
}

func (s *Sorter) sort(array []int) {
	s.strategy.execute()
}

func main() {
	sorter := &Sorter{}
	quickSort := &QuickSort{}
	insertionSort := &InsertionSort{}
	
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if len(array) > 5 {
		sorter.setStrategy(quickSort)
	} else {
		sorter.setStrategy(insertionSort)
	}

	sorter.sort(array)
}
