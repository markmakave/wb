package main

// facade pattern
// facade pattern is a structural design pattern that provides a simplified interface to a library, a framework, or any other complex set of classes.

import (
	"fmt"
)

type CPU struct {
}

func (c *CPU) freeze() {
	fmt.Println("CPU freeze")
}

func (c *CPU) jump(position int) {
	fmt.Println("CPU jump", position)
}

func (c *CPU) execute() {
	fmt.Println("CPU execute")
}

type Memory struct {
}

func (m *Memory) load(position int, data string) {
	fmt.Println("Memory load", position, data)
}

type HardDrive struct {
}

func (h *HardDrive) read(lba, size int) string {
	fmt.Println("HardDrive read", lba, size)
	return "HardDrive read"
}

type Computer struct {
	cpu    *CPU
	memory *Memory
	hard   *HardDrive
}

func (c *Computer) startComputer() {
	c.cpu.freeze()
	c.memory.load(0, "boot")
	c.hard.read(0, 1)
	c.cpu.jump(0)
	c.cpu.execute()
}

func main() {
	computer := &Computer{
		cpu:    &CPU{},
		memory: &Memory{},
		hard:   &HardDrive{},
	}
	computer.startComputer()
}
