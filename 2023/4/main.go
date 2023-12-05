package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/sek-consulting/aoc/lib"
)

func main() {
	// lib.Expect(part1("example1.txt"), "13")

	// result1 := part1("input.txt")
	// fmt.Println(result1)

	lib.Expect(part2("example2.txt"), "30")
}

func part1(filename string) string {
	sum := 0

	for _, line := range lib.ReadLines(filename) {
		split := strings.Split(line, ": ")
		split = strings.Split(split[1], " | ")

		winners := strings.Split(split[0], " ")
		numbers := strings.Split(split[1], " ")
		matches := countMatches(numbers, winners)

		if matches > 0 {
			sum += lib.Pow(2, matches-1)
		}

	}
	return strconv.Itoa(sum)
}

func countMatches(numbers, winners []string) int {
	matches := 0
	for _, num := range numbers {
		if lib.IsNum(num) && slices.Contains(winners, num) {
			matches += 1
		}
	}
	return matches
}

func part2(filename string) string {
	return ""
}
