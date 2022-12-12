package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("AoC Day 7 Part 2")

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

	scenicScore := 0
	for y, yv := range grid {
		for x, _ := range yv {
			score := checkVis(grid, x, y, false)
			if score > scenicScore {
				scenicScore = score
			}
		}
	}
	fmt.Println(scenicScore)
}

func checkVis(grid [][]int, x, y int, dbg bool) int {
	// Check North
	nScore := 0
	for i := y - 1; i >= 0; i-- {
		nScore++
		if grid[i][x] >= grid[y][x] {
			break
		}
	}
	// Check South
	sScore := 0
	for i := y + 1; i < len(grid); i++ {
		sScore++
		if grid[i][x] >= grid[y][x] {
			break
		}
	}

	// Check East
	eScore := 0
	for i := x + 1; i < len(grid[y]); i++ {
		eScore++
		if grid[y][i] >= grid[y][x] {
			break
		}
	}

	// Check West
	wScore := 0
	for i := x - 1; i >= 0; i-- {
		wScore++
		if grid[y][i] >= grid[y][x] {
			break
		}
	}

	return nScore * wScore * sScore * eScore
}
