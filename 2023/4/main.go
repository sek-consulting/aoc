package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/sek-consulting/aoc/lib"
)

func main() {
	// lib.Expect(part1("example1.txt"), "13")

	// result1 := part1("input.txt")
	// fmt.Println(result1)

	// lib.Expect(part2("example1.txt"), "30")

	result2 := part2("input.txt")
	fmt.Println(result2)
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

	var wins []int
	for _, line := range lib.ReadLines(filename) {
		split := strings.Split(line, ": ")
		split = strings.Split(split[1], " | ")

		winners := strings.Split(split[0], " ")
		numbers := strings.Split(split[1], " ")
		wins = append(wins, countMatches(numbers, winners))
	}
	// fmt.Println(wins)

	tickets := make([]int, len(wins))
	for i := range tickets {
		tickets[i] += 1

		// fmt.Println(i, "factor", tickets[i], "wins", wins[i])

		for j := 1; j <= wins[i]; j++ {
			if i+j < len(tickets) {
				// fmt.Println("set", i+j, tickets[i+j], "+", tickets[i])
				tickets[i+j] += tickets[i]
			}
		}

		// fmt.Println(tickets)
	}

	sum := lib.SumSlice(tickets)

	return strconv.Itoa(sum)
}
