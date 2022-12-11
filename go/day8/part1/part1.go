package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("AoC Day 7 Part 1")

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

	grid := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		for _, v := range line {
			height, _ := strconv.Atoi(string(v))
			row = append(row, height)
		}

		grid = append(grid, row)
	}

	visible := 0
	for y, yv := range grid {
		for x, _ := range yv {
			if checkVis(grid, x, y, false) {
				visible++
				if y == 0 || y == len(grid)-1 || x == 0 || x == len(grid[y])-1 {
				} else {
					// fmt.Printf("X:%d Y:%d V\n", x, y)
				}
			}
		}
	}
	fmt.Println(visible)
}

func checkVis(grid [][]int, x, y int, dbg bool) bool {
	if y == 0 || y == len(grid)-1 || x == 0 || x == len(grid[y])-1 {
		return true
	}

	// Check North
	for i := y - 1; i >= 0; i-- {
		if grid[i][x] >= grid[y][x] {
			break
		}

		if i == 0 {
			return true
		}
	}
	// Check South
	for i := y + 1; i < len(grid); i++ {
		if grid[i][x] >= grid[y][x] {
			break
		}

		if i == len(grid)-1 {
			return true
		}
	}

	// Check East
	for i := x + 1; i < len(grid[y]); i++ {
		if grid[y][i] >= grid[y][x] {
			break
		}

		if i == len(grid[y])-1 {
			return true
		}
	}

	// Check West
	for i := x - 1; i >= 0; i-- {
		if grid[y][i] >= grid[y][x] {
			break
		}

		if i == 0 {
			return true
		}
	}

	return false
}
