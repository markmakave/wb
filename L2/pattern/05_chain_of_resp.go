package main

import (
	"fmt"
)

type Handler interface {
	setNext(handler Handler)
	handle(request string)
}

type AbstractHandler struct {
	next Handler
}

func (a *AbstractHandler) setNext(handler Handler) {
	a.next = handler
}

func (a *AbstractHandler) handle(request string) {
	if a.next != nil {
		a.next.handle(request)
	}
}

type MonkeyHandler struct {
	AbstractHandler
}

func (m *MonkeyHandler) handle(request string) {
	if request == "Banana" {
		fmt.Println("Monkey: I'll eat the", request)
	} else {
		m.AbstractHandler.handle(request)
	}
}

type SquirrelHandler struct {
	AbstractHandler
}

func (s *SquirrelHandler) handle(request string) {
	if request == "Nut" {
		fmt.Println("Squirrel: I'll eat the", request)
	} else {
		s.AbstractHandler.handle(request)
	}
}

type DogHandler struct {
	AbstractHandler
}

func (d *DogHandler) handle(request string) {
	if request == "MeatBall" {
		fmt.Println("Dog: I'll eat the", request)
	} else {
		d.AbstractHandler.handle(request)
	}
}

func main() {
	monkey := &MonkeyHandler{}
	squirrel := &SquirrelHandler{}
	dog := &DogHandler{}

	monkey.setNext(squirrel)
	squirrel.setNext(dog)

	monkey.handle("Nut")
	monkey.handle("Banana")
	monkey.handle("Cup of coffee")
}
