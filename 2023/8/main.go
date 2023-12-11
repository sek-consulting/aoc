package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sek-consulting/aoc/lib"
)

func main() {
	// lib.Expect(part1("example.txt"), "6")
	// fmt.Println(part1("input.txt"))

	// lib.Expect(part2("example2.txt"), "6")
	fmt.Println(part2("input.txt"))
}

type Graph map[string]Node
type Node struct {
	left  string
	right string
}

func part1(filename string) string {
	lines := lib.ReadLines(filename)

	instructions := lines[0]

	graph := Graph{}
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		graph[line[0:3]] = Node{line[7:10], line[12:15]}
	}

	// fmt.Println(graph)

	steps := getStepCount(graph, "AAA", "ZZZ", instructions)

	return strconv.Itoa(steps)
}

func getStepCount(graph Graph, curNode, target, instructions string) int {
	fmt.Println(curNode, target)
	steps := 0
	for !strings.HasSuffix(curNode, target) {
		move := instructions[steps%len(instructions)]

		// fmt.Println("STEP", steps, string(move))

		if move == 'L' {
			curNode = graph[curNode].left
		} else {
			curNode = graph[curNode].right
		}

		steps += 1
	}
	fmt.Println("->", steps, curNode)
	return steps
}

func part2(filename string) string {
	lines := lib.ReadLines(filename)

	instructions := lines[0]

	graph := Graph{}
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		graph[line[0:3]] = Node{line[7:10], line[12:15]}
	}

	// fmt.Println(graph)

	var steps []int
	for name := range graph {
		if strings.HasSuffix(name, "A") {
			steps = append(steps, getStepCount(graph, name, "Z", instructions))
		}
	}

	fmt.Println(steps)

	result := lcm(steps[0], steps[1], steps[2:]...)

	return strconv.Itoa(result)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, nums ...int) int {
	result := a * b / gcd(a, b)
	for _, i := range nums {
		result = lcm(result, i)
	}
	return result
}
