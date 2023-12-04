package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sek-consulting/aoc/lib"
)

func main() {
	lib.Expect(part1("example1.txt"), "8")

	result1 := part1("input.txt")
	fmt.Println("part 1:", result1)

	// result2 := part2("input.txt")
	// fmt.Println("part 2:", result2)
}

func part2(filename string) {

}

func part1(filename string) string {
	max := make(map[string]int)
	max["red"] = 12
	max["blue"] = 14
	max["green"] = 13

	sum := 0

	for i, line := range lib.ReadLines(filename) {
		rounds := splitIntoRounds(line)
		isValid := true
		for _, game := range rounds {
			container := getBalls(game)
			if isOverMax(container, max) {
				isValid = false
			}
		}
		if isValid {
			sum += (i + 1)
		}
	}
	return strconv.Itoa(sum)
}

func splitIntoRounds(s string) []string {
	split := strings.Split(s, ":")
	rounds := strings.Split(split[1], ";")
	return rounds
}

func getBalls(s string) map[string]int {
	result := make(map[string]int)
	result["red"] = 0
	result["blue"] = 0
	result["green"] = 0

	balls := strings.Split(strings.TrimSpace(s), ",")
	for _, ball := range balls {
		v := strings.Split(strings.TrimSpace(ball), " ")
		result[v[1]] += lib.Atoi(v[0])
	}
	return result
}

func isOverMax(container, max map[string]int) bool {
	for color, num := range max {
		if container[color] > num {
			return true
		}
	}
	return false
}
