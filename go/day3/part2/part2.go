package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("AoC Day 3 Part 2")

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
		bag1 := scanner.Text()
		scanner.Scan()
		bag2 := scanner.Text()
		scanner.Scan()
		bag3 := scanner.Text()

		for i := 0; i < len(bag1); i++ {
			for n := 0; n < len(bag2); n++ {
				if bag1[i] == bag2[n] {
					for c := 0; c < len(bag3); c++ {
						if bag2[n] == bag3[c] && bag3[c] == bag1[i] {
							total_priority = total_priority + getPriority(bag1[i])
							goto end
						}
					}
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
