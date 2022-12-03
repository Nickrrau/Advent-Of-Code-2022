package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("AoC Day 1 Part 2")

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

	top_three_cals := [3]int{0, 0, 0}
	tmp_cal := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		if len(input) < 1 {
			pushPop(&top_three_cals, tmp_cal)
			tmp_cal = 0
		} else {
			cal, err := strconv.Atoi(input)
			if err != nil {
				break
			}
			tmp_cal = tmp_cal + cal
			continue
		}
	}
	pushPop(&top_three_cals, tmp_cal)

	fmt.Println(top_three_cals[0] + top_three_cals[1] + top_three_cals[2])
}

func pushPop(stack *[3]int, value int) {
	if stack[0] < value {
		stack[2] = stack[1]
		stack[1] = stack[0]
		stack[0] = value
	} else if stack[1] < value {
		stack[2] = stack[1]
		stack[1] = value
	} else if stack[2] < value {
		stack[2] = value
	}
}