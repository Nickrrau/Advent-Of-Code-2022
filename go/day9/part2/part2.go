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
	fmt.Println("AoC Day 9 Part 2")

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

func CreateSegment(x, y int) *RopeSegment {
	return &RopeSegment{Coord{0, 0}, []Coord{{0, 0}}}
}

type Rope struct {
	head     RopeSegment
	segments []RopeSegment
}

func CreateRope(numOfKnots int) *Rope {
	head := CreateSegment(0, 0)
	knots := []RopeSegment{}
	for i := 0; i < numOfKnots; i++ {
		knots = append(knots, *CreateSegment(0, 0))
	}
	return &Rope{
		*head,
		knots,
	}
}

func (r *Rope) moveSegments() {
	for s := 0; s < len(r.segments); s++ {
		var prev RopeSegment
		if s == 0 {
			prev = r.head
		} else {
			prev = r.segments[s-1]
		}
		if !r.segments[s].AdjacentTo(prev) {
			dX := prev.x - r.segments[s].x
			dY := prev.y - r.segments[s].y

			if math.Abs(float64(dX)) <= 2 && math.Abs(float64(dY)) <= 2 {
				if dX > 1 {
					dX = 1
				}
				if dX < -1 {
					dX = -1
				}
				if dY > 1 {
					dY = 1
				}
				if dY < -1 {
					dY = -1
				}
			} else if math.Abs(float64(dX)) == 2 && math.Abs(float64(dY)) == 0 {
				if dX > 1 {
					dX = 1
				}
				if dX < -1 {
					dX = -1
				}
			} else if math.Abs(float64(dX)) == 0 && math.Abs(float64(dY)) == 2 {
				if dY > 1 {
					dY = 1
				}
				if dY < -1 {
					dY = -1
				}
			}
			r.segments[s].move(r.segments[s].x+dX, r.segments[s].y+dY)
			r.segments[s].history = append(r.segments[s].history, Coord{r.segments[s].x, r.segments[s].y})
		} else {
			break
		}
	}
}

func (r *Rope) Move(direction string, distance int) {
	for i := 0; i < distance; i++ {
		switch direction {
		case "R":
			r.head.move(r.head.x+1, r.head.y)
			r.moveSegments()
			break
		case "L":
			r.head.move(r.head.x-1, r.head.y)
			r.moveSegments()
			break
		case "U":
			r.head.move(r.head.x, r.head.y+1)
			r.moveSegments()
			break
		case "D":
			r.head.move(r.head.x, r.head.y-1)
			r.moveSegments()
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

	rope := CreateRope(9)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		distance, _ := strconv.Atoi(line[1])
		rope.Move(line[0], distance)
	}

	uniqueSteps := []Coord{}
	for _, v := range rope.segments[len(rope.segments)-1].history {
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
