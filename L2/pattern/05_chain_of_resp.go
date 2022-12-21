package main

// chain of responsibility pattern
// chain of responsibility pattern is a behavioral design pattern that lets you pass requests along a chain of handlers. Upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain.
// The chain of responsibility pattern is a way to decouple senders and receivers of requests based on type of request. This pattern is essentially a linear search for an object that can handle a particular request.
// The chain of responsibility pattern lets you:
// - pass requests along a chain of handlers
// - upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain
// - optionally, each handler may decide to process the request and stop passing it further along the chain
// - the pattern lets you modify the chain of handlers at runtime
// - the pattern lets you decouple senders and receivers of requests based on type of request

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
