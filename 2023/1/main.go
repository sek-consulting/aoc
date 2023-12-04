package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/sek-consulting/aoc/lib"
)

func main() {
	lib.Expect(part2("example2.txt"), "281")

	result1 := part1("input.txt")
	fmt.Println("part 1:", result1)

	result2 := part2("input.txt")
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

func part2(filename string) string {
	sum := 0

	rgxFirst := regexp.MustCompile(".*?(?P<num>[0-9]|one|two|three|four|five|six|seven|eight|nine).*")
	rgxLast := regexp.MustCompile(".*(?P<num>[0-9]|one|two|three|four|five|six|seven|eight|nine).*?")

	for _, line := range lib.ReadLines(filename) {
		first := rgxFirst.FindStringSubmatch(line)[1]
		first = convertToNumbers(first)

		last := rgxLast.FindStringSubmatch(line)[1]
		last = convertToNumbers(last)

		numStr := first + last
		num := lib.Atoi(numStr)

		sum += num
	}
	return strconv.Itoa(sum)
}

func convertToNumbers(s string) string {
	lookup := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i, val := range lookup {
		s = strings.ReplaceAll(s, val, strconv.Itoa(i))
	}
	return s
}

func removeChars(s string) string {
	rgx := regexp.MustCompile("[^0-9]")
	return rgx.ReplaceAllString(s, "")
}
