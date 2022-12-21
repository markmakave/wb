package main

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
