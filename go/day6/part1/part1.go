package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("AoC Day 6 Part 1")

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

	pos := 3

	scanner := bufio.NewScanner(file)
	scanner.Split(signalSplitFunc)
	for scanner.Scan() {
		line := scanner.Bytes()

		pos = pos + 1

		signal := map[byte]int{}
		signal[line[0]] = 1
		signal[line[1]] = 1
		signal[line[2]] = 1
		signal[line[3]] = 1

		if len(signal) == 4 {
			break
		}
	}

	fmt.Println(pos)
}

func signalSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if len(data) < 4 {
		return len(data), data, nil
	}

	return 1, data[:4], nil
}
