package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("AoC Day 6 Part 2")

	if len(os.Args) < 2 {
		fmt.Println("Need to provide input")
		return
	}

	input := os.Args[1]
	processInput(input)
}

func processInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pos := 13

	scanner := bufio.NewScanner(file)
	scanner.Split(signalSplitFunc)
	for scanner.Scan() {
		line := scanner.Bytes()

		pos = pos + 1

		signal := map[byte]int{}
		for _, v := range line {
			signal[v] = 1
		}

		if len(signal) == 14 {
			break
		}
	}

	fmt.Println(pos)
}

func signalSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if len(data) < 14 {
		return len(data), data, nil
	}

	return 1, data[:14], nil
}
