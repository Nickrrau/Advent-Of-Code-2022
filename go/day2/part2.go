package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("AoC Day 2 Part 2")

	if len(os.Args) < 2 {
		fmt.Println("Need to provide input")
		return
	}

	input := os.Args[1]
	processInput(input)
}

type Outcome = byte
type Move = int

const (
	win  Outcome = 'Z'
	lose Outcome = 'X'
	draw Outcome = 'Y'
)
const (
	rock     Move = 1
	paper    Move = 2
	scissors Move = 3
)

func processInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	total_score := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		set := scanner.Text()
		if len(set) < 3 {
			panic("Invalid command set\n")
		}
		opponent_move := set[0]
		result := set[2]

		total_score = total_score + processMoves(result, byteToMove(opponent_move))
	}

	fmt.Println(total_score)
}

func byteToMove(move byte) Move {
	if move == 'X' || move == 'A' {
		return rock
	}
	if move == 'Y' || move == 'B' {
		return paper
	}
	if move == 'Z' || move == 'C' {
		return scissors
	}
	return 0
}

func processMoves(result Outcome, opponent Move) int {
	if result == draw {
		return opponent + 3
	}

	if result == lose {
		switch opponent {
		case rock:
			return scissors
		case paper:
			return rock
		case scissors:
			return paper
		}
	}

	switch opponent {
	case rock:
		return paper + 6
	case paper:
		return scissors + 6
	case scissors:
		return rock + 6
	}

	return 0
}
