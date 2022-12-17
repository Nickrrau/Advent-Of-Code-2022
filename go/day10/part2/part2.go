package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("AoC Day 10 Part 2")

	if len(os.Args) < 2 {
		fmt.Println("Need to provide input")
		return
	}

	input := os.Args[1]
	processInput(input)
}

type Instruction int

const (
	NOOP Instruction = 0
	ADDX Instruction = 1
)

type InstructionCmd struct {
	t Instruction
	v int
}

type Emulator struct {
	signalRegister int
	instruction    InstructionCmd
}

func NewEmulator() *Emulator {
	return &Emulator{
		signalRegister: 1,
		instruction:    InstructionCmd{NOOP, 0},
	}
}

func (e *Emulator) LoadInstruction(i Instruction, v int) {
	e.instruction = InstructionCmd{i, v}
}

func (e *Emulator) Cycle() {
	switch e.instruction.t {
	case NOOP:
		break
	case ADDX:
		e.signalRegister = e.signalRegister + e.instruction.v
		break
	}
}

type CRT struct {
	position int
	rows     []bool
}

func NewCRT() *CRT {
	return &CRT{
		0,
		[]bool{},
	}
}

func (c *CRT) Cycle(spritePos int) {
	row := (c.position) / 40
	if ((c.position)-(row*40)) >= spritePos-1 && ((c.position)-(row*40)) <= spritePos+1 {
		c.rows = append(c.rows, true)
	} else {
		c.rows = append(c.rows, false)
	}
	c.position++
}

func processInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	cpu := NewEmulator()
	crt := NewCRT()
	crt.Cycle(cpu.signalRegister) // dumb trick

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parsed := bytes.Split(scanner.Bytes(), []byte{' '})
		switch string(parsed[0]) {
		case "noop":
			cpu.LoadInstruction(NOOP, 0)
			cpu.Cycle()
			crt.Cycle(cpu.signalRegister)
			break
		case "addx":
			cpu.LoadInstruction(NOOP, 0)
			cpu.Cycle()
			crt.Cycle(cpu.signalRegister)
			v, _ := strconv.Atoi(string(parsed[1]))
			cpu.LoadInstruction(ADDX, v)
			cpu.Cycle()
			crt.Cycle(cpu.signalRegister)
			break
		}
	}

	for i, v := range crt.rows {
		if i%40 == 0 && i != 0 {
			fmt.Println()
		}
		if v {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
}
