package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("AoC Day 9 Part 1")

	if len(os.Args) < 2 {
		fmt.Println("Need to provide input")
		return
	}

	input := os.Args[1]
	processInput(input)
}

type Coord struct {
	x, y int
}

type RopeSegment struct {
	Coord
	history []Coord
}

func (r *RopeSegment) move(x, y int) {
	r.x = x
	r.y = y
}

func (r *RopeSegment) AdjacentTo(segment RopeSegment) bool {
	return math.Trunc(math.Sqrt((math.Pow(float64(segment.x-r.x), 2))+(math.Pow(float64(segment.y-r.y), 2)))) < 2
}

func CreateSegment(x, y int) RopeSegment {
	return RopeSegment{Coord{0, 0}, []Coord{{0, 0}}}
}

type Rope struct {
	head RopeSegment
	tail RopeSegment
}

func (r *Rope) Move(direction string, distance int) {
	for i := 0; i < distance; i++ {
		switch direction {
		case "R":
			r.head.move(r.head.x+1, r.head.y)
			if !r.tail.AdjacentTo(r.head) {
				r.tail.move(r.head.x-1, r.head.y)
				r.tail.history = append(r.tail.history, Coord{r.tail.x, r.tail.y})
			}
			break
		case "L":
			r.head.move(r.head.x-1, r.head.y)
			if !r.tail.AdjacentTo(r.head) {
				r.tail.move(r.head.x+1, r.head.y)
				r.tail.history = append(r.tail.history, Coord{r.tail.x, r.tail.y})
			}
			break
		case "U":
			r.head.move(r.head.x, r.head.y+1)
			if !r.tail.AdjacentTo(r.head) {
				r.tail.move(r.head.x, r.head.y-1)
				r.tail.history = append(r.tail.history, Coord{r.tail.x, r.tail.y})
			}
			break
		case "D":
			r.head.move(r.head.x, r.head.y-1)
			if !r.tail.AdjacentTo(r.head) {
				r.tail.move(r.head.x, r.head.y+1)
				r.tail.history = append(r.tail.history, Coord{r.tail.x, r.tail.y})
			}
			break
		}
		r.head.history = append(r.head.history, Coord{r.head.x, r.head.y})
	}
}

func processInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rope := &Rope{head: CreateSegment(0, 0), tail: CreateSegment(0, 0)}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		distance, _ := strconv.Atoi(line[1])
		rope.Move(line[0], distance)
	}

	uniqueSteps := []Coord{}
	for _, v := range rope.tail.history {
		for i := 0; i < len(uniqueSteps); i++ {
			if v == uniqueSteps[i] {
				goto skip
			}
		}
		uniqueSteps = append(uniqueSteps, v)
	skip:
	}
	fmt.Println(len(uniqueSteps))
}
