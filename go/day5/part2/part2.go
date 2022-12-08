package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("AoC Day 5 Part 2")

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

	stack := make(map[byte][]byte, 1)
	order := []byte{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Bytes()) < 1 {
			break
		}
		stack[scanner.Bytes()[0]] = bytes.ReplaceAll(scanner.Bytes()[1:], []byte{' '}, []byte{})
		order = append(order, scanner.Bytes()[0])
	}

	for scanner.Scan() {
		line := bytes.Split(scanner.Bytes(), []byte{' '})
		move(&stack, line[1], line[3][0], line[5][0])
	}

	for _, v := range order {
		fmt.Printf("%s", string(stack[v][len(stack[v])-1]))
	}
}

func move(stacks *map[byte][]byte, count []byte, originStack byte, targetStack byte) {
	num, _ := strconv.Atoi(string(count))
	originItems := (*stacks)[originStack][len((*stacks)[originStack])-num:]
	targetItems := (*stacks)[targetStack]

	// for i := len(originItems) - 1; i > -1; i-- { // Nice...
	targetItems = append(targetItems, originItems...)
	// }

	(*stacks)[originStack] = (*stacks)[originStack][0 : len((*stacks)[originStack])-num]
	(*stacks)[targetStack] = targetItems
}
