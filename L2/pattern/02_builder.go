package main

// builder pattern
// builder pattern is a creational design pattern that lets you construct complex objects step by step. The pattern allows you to produce different types and representations of an object using the same construction code.
// The Builder pattern is a good choice when designing classes whose constructors or static factories would have more than a handful of parameters.

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

type ComputerBuilder struct {
	cpu    *CPU
	memory *Memory
	hard   *HardDrive
}

func (c *ComputerBuilder) setCPU() {
	c.cpu = &CPU{}
}

func (c *ComputerBuilder) setMemory() {
	c.memory = &Memory{}
}

func (c *ComputerBuilder) setHardDrive() {
	c.hard = &HardDrive{}
}

func (c *ComputerBuilder) getComputer() *Computer {
	return &Computer{
		cpu:    c.cpu,
		memory: c.memory,
		hard:   c.hard,
	}
}

func main() {
	computerBuilder := &ComputerBuilder{}
	computerBuilder.setCPU()
	computerBuilder.setMemory()
	computerBuilder.setHardDrive()
	computer := computerBuilder.getComputer()
	computer.startComputer()
}
