package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/sek-consulting/aoc/lib"
)

func main() {
	result1 := part1("./2023/1/input.txt")
	fmt.Println("part 1:", result1)

	result2 := part2("./2023/1/input.txt")
	fmt.Println("part 2:", result2)
}

func part1(filename string) int {
	sum := 0
	for _, line := range lib.ReadLines(filename) {
		onlyNums := removeChars(line)
		numsArr := strings.Split(onlyNums, "")
		numStr := numsArr[0] + numsArr[len(numsArr)-1]
		num := lib.Atoi(numStr)
		sum += num
	}
	return sum
}

func part2(filename string) int {
	sum := 0
	for _, line := range lib.ReadLines(filename) {
		fmt.Println(line)
	}
	return sum
}

func removeChars(s string) string {
	rgx := regexp.MustCompile("[^0-9]")
	return rgx.ReplaceAllString(s, "")
}
