package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Println("AoC Day 12 Part 1")

	if len(os.Args) < 2 {
		fmt.Println("Need to provide input")
		return
	}

	input := os.Args[1]
	processInput(input)
}

type Cell struct {
	grid   *Grid
	x, y   int
	height int
	prev   *Cell
}

func (c Cell) String() string {
	return fmt.Sprintf("%c(%d,%d)", c.height, c.x, c.y)
}

func (c *Cell) GetNeighbors() []*Cell {
	cells := []*Cell{}

	if c.x > 0 {
		cells = append(cells, &c.grid.values[c.y][c.x-1])
	}
	if c.x < len(c.grid.values[c.y])-1 {
		cells = append(cells, &c.grid.values[c.y][c.x+1])
	}
	if c.y > 0 {
		cells = append(cells, &c.grid.values[c.y-1][c.x])
	}
	if c.y < len(c.grid.values)-1 {
		cells = append(cells, &c.grid.values[c.y+1][c.x])
	}

	return cells
}

type Grid struct {
	pos    struct{ x, y int }
	end    struct{ x, y int }
	values [][]Cell
}

func (g *Grid) SetPos(x, y int) {
	g.pos.x = x
	g.pos.y = y
}
func (g *Grid) SetEnd(x, y int) {
	g.end.x = x
	g.end.y = y
}

func (g *Grid) InsertCell(x, y int, height byte) {
	switch height {
	case 'S':
		height = 'a'
		g.SetPos(x, y)
		break
	case 'E':
		height = 'z'
		g.SetEnd(x, y)
		break
	}
	if len(g.values)-1 < y {
		g.values = append(g.values, []Cell{})
	}
	g.values[y] = append(g.values[y], Cell{g, x, y, int(height), nil})
}

func processInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	graph := Grid{values: [][]Cell{{}}}

	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		for x := 0; x < len(scanner.Bytes()); x++ {
			graph.InsertCell(x, y, scanner.Bytes()[x])
		}
	}

	fmt.Println(FindPath(
		&graph.values[graph.end.y][graph.end.x],
		&graph.values[graph.pos.y][graph.pos.x],
	))
}

func FindPath(target, start *Cell) (bool, int) {
	visited := []*Cell{}
	queue := []*Cell{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == target {
			break
		}

		neighbors := current.GetNeighbors()
		for c := 0; c < len(neighbors); c++ {
			if math.Abs(float64(neighbors[c].height-current.height)) < 2 || neighbors[c].height <= current.height {
				if neighbors[c].height >= current.height || neighbors[c].height < current.height {
					visMatch := false
					for i := 0; i < len(visited); i++ {
						if neighbors[c] == visited[i] {
							visMatch = true
							break
						}
					}

					if !visMatch {
						visited = append(visited, neighbors[c])
						neighbors[c].prev = current
						queue = append(queue, neighbors[c])
					}
				}
			}
		}
	}

	length := 0
	prev := target
	for prev != start {
		if prev.prev != nil {
			length++
			prev = prev.prev
		} else {
			break
		}
	}

	return true, length
}
