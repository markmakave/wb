package main

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

type ComputerPartVisitor interface {
	visitCPU(cpu *CPU)
	visitMemory(memory *Memory)
	visitHardDrive(hardDrive *HardDrive)
}

type ComputerPartDisplayVisitor struct {
}

func (c *ComputerPartDisplayVisitor) visitCPU(cpu *CPU) {
	fmt.Println("Displaying CPU")
}

func (c *ComputerPartDisplayVisitor) visitMemory(memory *Memory) {
	fmt.Println("Displaying Memory")
}

func (c *ComputerPartDisplayVisitor) visitHardDrive(hardDrive *HardDrive) {
	fmt.Println("Displaying HardDrive")
}

type ComputerPart interface {
	accept(visitor ComputerPartVisitor)
}

func (c *CPU) accept(visitor ComputerPartVisitor) {
	visitor.visitCPU(c)
}

func (m *Memory) accept(visitor ComputerPartVisitor) {
	visitor.visitMemory(m)
}

func (h *HardDrive) accept(visitor ComputerPartVisitor) {
	visitor.visitHardDrive(h)
}

func (c *Computer) accept(visitor ComputerPartVisitor) {
	for _, part := range []ComputerPart{c.cpu, c.memory, c.hard} {
		part.accept(visitor)
	}
	visitor.visitCPU(c.cpu)
}

func main() {
	computer := &Computer{
		cpu:    &CPU{},
		memory: &Memory{},
		hard:   &HardDrive{},
	}
	computer.startComputer()

	computer.accept(&ComputerPartDisplayVisitor{})
}
