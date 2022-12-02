package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("AoC Day 1 Part 1")

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

	highest_cal := 0
	tmp_cal := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			tmp_cal = 0
			continue
		}

		cal, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		tmp_cal = tmp_cal + cal

		if highest_cal < tmp_cal {
			highest_cal = tmp_cal
		}
	}
	fmt.Println(highest_cal)
}