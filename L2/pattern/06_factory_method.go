package main

// factory method pattern
// factory method pattern is a creational design pattern that provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created.
// The factory method pattern solves problems like:
// - How can an object be created so that subclasses can redefine which class to instantiate?
// - How can a class defer instantiation to subclasses?

import (
	"fmt"
)

type Computer interface {
	start()
}

type Mac struct {
}

func (m *Mac) start() {
	fmt.Println("Mac start")
}

type Windows struct {
}

func (w *Windows) start() {
	fmt.Println("Windows start")
}

type ComputerFactory struct {
}

func (c *ComputerFactory) getComputer(computerType string) Computer {
	if computerType == "mac" {
		return &Mac{}
	} else if computerType == "windows" {
		return &Windows{}
	}
	return nil
}

func main() {
	computerFactory := &ComputerFactory{}
	mac := computerFactory.getComputer("mac")
	mac.start()
	windows := computerFactory.getComputer("windows")
	windows.start()
}
