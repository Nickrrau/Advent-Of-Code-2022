package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("AoC Day 2 Part 1")

	if len(os.Args) < 2 {
		fmt.Println("Need to provide input")
		return
	}

	input := os.Args[1]
	processInput(input)
}

type Move = int

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
		player_move := set[2]

		total_score = total_score + processMoves(byteToMove(player_move), byteToMove(opponent_move))
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

func processMoves(player Move, opponent Move) int {
	if player == opponent {
		return player + 3
	}

	if (player == rock && opponent == paper) || (player == paper && opponent == scissors) || (player == scissors && opponent == rock) {
		return player
	}

	return player + 6
}