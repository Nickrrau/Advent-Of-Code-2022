package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("AoC Day 3 Part 1")

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

	total_priority := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		comp1 := line[0 : len(line)/2]
		comp2 := line[len(line)/2:]

		for i := 0; i < len(comp1); i++ {
			for n := 0; n < len(comp2); n++ {
				if comp1[i] == comp2[n] {
					total_priority = total_priority + getPriority(comp1[i])
					goto end
				}
			}
		}
	end:
	}

	fmt.Println(total_priority)
}

func getPriority(v byte) int {
	if v > 64 && v < 91 {
		return 52 - (90 - int(v))
	}
	if v > 96 && v < 123 {
		return 26 - (122 - int(v))
	}
	return 0
}
