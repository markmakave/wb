package main

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
