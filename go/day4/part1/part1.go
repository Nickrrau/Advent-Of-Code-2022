package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("AoC Day 4 Part 1")

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

	overlaps := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ranges := strings.Split(scanner.Text(), ",")
		if isOverlaping(strings.Split(ranges[0], "-"), strings.Split(ranges[1], "-")) {
			overlaps++
		}
	}

	fmt.Println(overlaps)
}

func isOverlaping(range1 []string, range2 []string) bool {
	beg1, _ := strconv.ParseInt(range1[0], 10, 64)
	beg2, _ := strconv.ParseInt(range2[0], 10, 64)

	end1, _ := strconv.ParseInt(range1[1], 10, 64)
	end2, _ := strconv.ParseInt(range2[1], 10, 64)

	if beg1 == beg2 && end1 == end2 {
		return true
	}

	// r1:        beg1 ---- end1
	// r2:   beg2 ---- end2
	if beg1 > beg2 && end1 > end2 {
		return false
	}

	// r1:  beg1 ---- end1
	// r2:       beg2 ---- end2
	if beg1 < beg2 && end1 < end2 {
		return false
	}

	return true
}