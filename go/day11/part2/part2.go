package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("AoC Day 11 Part 2")

	if len(os.Args) < 2 {
		fmt.Println("Need to provide input")
		return
	}

	input := os.Args[1]
	processInput(input)
}

type OperationType int

const (
	MULT OperationType = 0
	ADD  OperationType = 1
	SUB  OperationType = 2
	DIV  OperationType = 3
	MOD  OperationType = 4
)

type OperationValue struct {
	old bool
	val int
}

type Operation struct {
	kind   OperationType
	v1, v2 OperationValue
}

type Monkey struct {
	items       []int
	op          Operation
	mod         int
	div         int
	trueTarget  int
	falseTarget int

	inspected int
}

func (m *Monkey) Inspect() {
	value1 := m.op.v2.val
	if m.op.v1.old {
		value1 = m.items[0]
	}
	value2 := m.op.v2.val
	if m.op.v2.old {
		value2 = m.items[0]
	}
	switch m.op.kind {
	case DIV:
		m.items[0] = value1 / value2
		break
	case MULT:
		m.items[0] = value1 * value2
		break
	case ADD:
		m.items[0] = value1 + value2
		break
	case SUB:
		m.items[0] = value1 - value2
		break
	}

	m.items[0] = m.items[0] % m.mod
	m.inspected++
}

func (m *Monkey) Test() int {
	if m.items[0]%m.div == 0 {
		return m.trueTarget
	}
	return m.falseTarget
}

func (m *Monkey) Throw(target *Monkey) {
	target.Catch(m.items[0])
	m.items = m.items[1:]
}

func (m *Monkey) Catch(item int) {
	m.items = append(m.items, item)
}

func processInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	monkeys := []Monkey{}
	mod := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scanner.Scan()
		items := strings.Split(scanner.Text(), ":")[1]
		scanner.Scan()
		op := strings.Split(scanner.Text(), ":")[1]
		scanner.Scan()

		test := strings.Split(scanner.Text(), ":")[1]
		scanner.Scan()
		testTrue := strings.Split(scanner.Text(), ":")[1]
		scanner.Scan()
		testFalse := strings.Split(scanner.Text(), ":")[1]
		scanner.Scan()

		div, tTrg, fTrg := parseTest(test, testTrue, testFalse)
		if mod == 0 {
			mod = div
		}
		mod = mod * div

		monkeys = append(monkeys, Monkey{
			items:       parseItems(items),
			op:          parseOp(op),
			div:         div,
			trueTarget:  tTrg,
			falseTarget: fTrg,
			inspected:   0,
		})
	}

	for m := 0; m < len(monkeys); m++ {
		monkeys[m].mod = mod
	}

	for i := 0; i < 10000; i++ {
		for m := 0; m < len(monkeys); m++ {
			numItems := len(monkeys[m].items)
			for c := 0; c < numItems; c++ {
				monkeys[m].Inspect()
				target := monkeys[m].Test()
				monkeys[m].Throw(&monkeys[target])
			}
		}
	}

	topTwo := [2]int{0, 0}
	for i := 0; i < len(monkeys); i++ {
		if monkeys[i].inspected > topTwo[0] {
			topTwo[1] = topTwo[0]
			topTwo[0] = monkeys[i].inspected
		} else if monkeys[i].inspected > topTwo[1] {
			topTwo[1] = monkeys[i].inspected
		}
	}

	fmt.Println(topTwo[0] * topTwo[1])
}

func parseItems(input string) []int {
	s_items := strings.Split(strings.ReplaceAll(input, " ", ""), ",")

	items := []int{}
	for _, v := range s_items {
		n, _ := strconv.Atoi(v)
		items = append(items, n)
	}

	return items
}

func parseOp(input string) Operation {
	tokens := strings.Split(strings.TrimSpace(input), " ")

	var v1 OperationValue
	if tokens[2] == "old" {
		v1.old = true
	} else {
		v, _ := strconv.Atoi(tokens[2])
		v1.old = false
		v1.val = v
	}

	var v2 OperationValue
	if tokens[4] == "old" {
		v2.old = true
	} else {
		v, _ := strconv.Atoi(tokens[4])
		v2.old = false
		v2.val = v
	}

	var op Operation
	switch tokens[3] {
	case "+":
		op.kind = ADD
		break
	case "*":
		op.kind = MULT
		break
	}
	op.v1 = v1
	op.v2 = v2

	return op
}

func parseTest(test, trueTest, falseTest string) (int, int, int) {
	div, _ := strconv.Atoi(strings.Split(strings.TrimSpace(test), " ")[2])
	tID, _ := strconv.Atoi(strings.Split(strings.TrimSpace(trueTest), " ")[3])
	fID, _ := strconv.Atoi(strings.Split(strings.TrimSpace(falseTest), " ")[3])

	return div, tID, fID
}