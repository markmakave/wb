package main

import (
	"fmt"
)

type Command interface {
	execute()
}

type Computer struct {
}

func (c *Computer) start() {
	fmt.Println("Computer start")
}

func (c *Computer) restart() {
	fmt.Println("Computer restart")
}

func (c *Computer) stop() {
	fmt.Println("Computer stop")
}

type StartCommand struct {
	computer *Computer
}

func (s *StartCommand) execute() {
	s.computer.start()
}

type RestartCommand struct {
	computer *Computer
}

func (r *RestartCommand) execute() {
	r.computer.restart()
}

type StopCommand struct {
	computer *Computer
}

func (s *StopCommand) execute() {
	s.computer.stop()
}

type Invoker struct {
	startCommand    Command
	restartCommand  Command
	stopCommand     Command
	startCommandLog []Command
}

func (i *Invoker) setStartCommand(startCommand Command) {
	i.startCommand = startCommand
}

func (i *Invoker) setRestartCommand(restartCommand Command) {
	i.restartCommand = restartCommand
}

func (i *Invoker) setStopCommand(stopCommand Command) {
	i.stopCommand = stopCommand
}

func (i *Invoker) startComputer() {
	i.startCommand.execute()
	i.startCommandLog = append(i.startCommandLog, i.startCommand)
}

func (i *Invoker) restartComputer() {
	i.restartCommand.execute()
	i.startCommandLog = append(i.startCommandLog, i.restartCommand)
}

func (i *Invoker) stopComputer() {
	i.stopCommand.execute()
	i.startCommandLog = append(i.startCommandLog, i.stopCommand)
}

func (i *Invoker) showStartCommandLog() {
	for _, command := range i.startCommandLog {
		command.execute()
	}
}

func main() {
	computer := &Computer{}
	startCommand := &StartCommand{computer: computer}
	restartCommand := &RestartCommand{computer: computer}
	stopCommand := &StopCommand{computer: computer}
	invoker := &Invoker{}
	invoker.setStartCommand(startCommand)
	invoker.setRestartCommand(restartCommand)
	invoker.setStopCommand(stopCommand)
	invoker.startComputer()
	invoker.restartComputer()
	invoker.stopComputer()
	invoker.showStartCommandLog()
}
