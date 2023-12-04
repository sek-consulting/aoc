package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sek-consulting/aoc/lib"
)

func main() {
	// lib.Expect(part1("example1.txt"), "8")
	// lib.Expect(part2("example1.txt"), "2286")

	result1 := part1("input.txt")
	fmt.Println("part 1:", result1)

	result2 := part2("input.txt")
	fmt.Println("part 2:", result2)
}

func part2(filename string) string {
	sum := 0
	for _, line := range lib.ReadLines(filename) {
		container := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		rounds := splitIntoRounds(line)
		for _, round := range rounds {
			for _, draw := range strings.Split(round, ", ") {
				v := strings.Split(draw, " ")
				container[v[1]] = max(container[v[1]], lib.Atoi(v[0]))
			}
		}
		product := 1
		for _, i := range container {
			product *= i
		}
		sum += product
	}
	return strconv.Itoa(sum)
}

func part1(filename string) string {
	max := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sum := 0
	for i, line := range lib.ReadLines(filename) {
		rounds := splitIntoRounds(line)
		isValid := true
		for _, round := range rounds {
			isValid = lib.Every[string](
				strings.Split(round, ", "),
				func(val string) bool {
					v := strings.Split(val, " ")
					return max[v[1]] >= lib.Atoi(v[0])
				})
			if !isValid {
				break
			}
		}
		fmt.Println(isValid)
		if isValid {
			sum += i + 1
		}
	}
	return strconv.Itoa(sum)
}

func splitIntoRounds(s string) []string {
	rounds := strings.Split(s, ": ")
	rounds = strings.Split(rounds[1], "; ")
	return rounds
}
