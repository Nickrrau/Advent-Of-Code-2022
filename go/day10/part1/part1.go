package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("AoC Day 10 Part 1")

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
	cycleCounter   int

	instruction InstructionCmd
}

func (e Emulator) String() string {
	return fmt.Sprintf(
		"X: %d\nC: %d\nS: %d\n",
		e.signalRegister,
		e.cycleCounter,
		e.GetSignalStrength(),
	)
}

func NewEmulator() *Emulator {
	return &Emulator{
		signalRegister: 1,
		cycleCounter:   1,
		instruction:    InstructionCmd{NOOP, 0},
	}
}

func (e Emulator) GetSignalStrength() int {
	return e.cycleCounter * e.signalRegister
}

func (e *Emulator) LoadInstruction(i Instruction, v int) {
	e.instruction = InstructionCmd{i, v}
}

func (e *Emulator) Cycle() {
	switch e.instruction.t {
	case NOOP:
		e.cycleCounter++
		break
	case ADDX:
		e.signalRegister = e.signalRegister + e.instruction.v
		e.cycleCounter++
		break
	}
}

func processInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	cpu := NewEmulator()
	signalStrength := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parsed := bytes.Split(scanner.Bytes(), []byte{' '})
		switch string(parsed[0]) {
		case "noop":
			cpu.LoadInstruction(NOOP, 0)
			cpu.Cycle()
			signalStrength = signalStrength + CheckStrength(*cpu)
			break
		case "addx":
			cpu.LoadInstruction(NOOP, 0)
			cpu.Cycle()
			signalStrength = signalStrength + CheckStrength(*cpu)
			v, _ := strconv.Atoi(string(parsed[1]))
			cpu.LoadInstruction(ADDX, v)
			cpu.Cycle()
			signalStrength = signalStrength + CheckStrength(*cpu)
			break
		}

	}

	fmt.Println(signalStrength)
}

func CheckStrength(e Emulator) int {
	if e.cycleCounter == 20 || (e.cycleCounter-20)%40 == 0 {
		return e.GetSignalStrength()
	}
	return 0
}
