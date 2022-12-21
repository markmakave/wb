package main

// state pattern
// state pattern is a behavioral design pattern that lets an object alter its behavior when its internal state changes. It appears as if the object changed its class.
// The state pattern lets you:
// - change the behavior of an object when its internal state changes
// - alter an object's behavior at run-time as a consequence of a state change
// - localize state-specific behavior and avoid massive conditional statements
// - implement state transitions explicitly in the context class
// - implement state transitions implicitly by defining state subclasses
// - define state objects that can be shared by multiple contexts
// - implement state machines in an object-oriented way
// - avoid subclass explosion
// - avoid the need for a large number of conditionals
// - avoid code duplication
// - avoid the need for a large number of subclasses
// - avoid the need for a large number of state objects
// - avoid the need for a large number of state subclasses

import (
	"fmt"
)

type State interface {
	execute()
}

type StartState struct {
}

func (s *StartState) execute() {
	fmt.Println("Computer start")
}

type RestartState struct {
}

func (r *RestartState) execute() {
	fmt.Println("Computer restart")
}

type StopState struct {
}

func (s *StopState) execute() {
	fmt.Println("Computer stop")
}

type Computer struct {
	state State
}

func (c *Computer) setState(state State) {
	c.state = state
}

func (c *Computer) executeState() {
	c.state.execute()
}

func main() {
	computer := &Computer{}
	startState := &StartState{}
	restartState := &RestartState{}
	stopState := &StopState{}
	computer.setState(startState)
	computer.executeState()
	computer.setState(restartState)
	computer.executeState()
	computer.setState(stopState)
	computer.executeState()
}
