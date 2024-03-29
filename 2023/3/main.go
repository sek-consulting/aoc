package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/sek-consulting/aoc/lib"
)

func main() {
	// lib.Expect(part1("example1.txt"), "4361")

	// result1 := part1("input.txt")
	// fmt.Println(result1)

	// lib.Expect(part2("example1.txt"), "467835")

	result2 := part2("input.txt")
	fmt.Println(result2)
}

type coord struct {
	y, x int
}

type number struct {
	num    int
	coords []coord
}

func part1(filename string) string {
	var num string = ""
	var numCoords, symbolCoords []coord
	var numbers []number

	for y, line := range lib.ReadLines(filename) {
		for x, char := range strings.Split(line, "") {

			if lib.IsNum(char) {
				if num == "" {
					numCoords = make([]coord, 0)
				}
				numCoords = append(numCoords, coord{y, x})
				num += char
			} else if num != "" {
				numbers = append(numbers, number{lib.Atoi(num), numCoords})
				num = ""
			}

			if isSymbol(char) {
				symbolCoords = append(symbolCoords, coord{y, x})
			}

		}
	}

	sum := 0
	for _, num := range numbers {
		if hasAdj(num.coords, symbolCoords) {
			sum += num.num
		}
	}

	return strconv.Itoa(sum)
}

func isSymbol(s string) bool {
	rgx := regexp.MustCompile("[^0-9.]")
	return rgx.MatchString(s)
}

func isAdj(a, b coord) bool {
	return lib.Abs(a.x-b.x) <= 1 && lib.Abs(a.y-b.y) <= 1
}

func hasAdj(a, b []coord) bool {
	for _, aCoord := range a {
		for _, bCoord := range b {
			if isAdj(aCoord, bCoord) {
				return true
			}
		}
	}
	return false
}

func part2(filename string) string {
	var num string = ""
	var numCoords, gearCoords []coord
	var numbers []number

	for y, line := range lib.ReadLines(filename) {
		for x, char := range strings.Split(line, "") {

			if lib.IsNum(char) {
				if num == "" {
					numCoords = make([]coord, 0)
				}
				numCoords = append(numCoords, coord{y, x})
				num += char
			} else if num != "" {
				numbers = append(numbers, number{lib.Atoi(num), numCoords})
				num = ""
			}

			if char == "*" {
				gearCoords = append(gearCoords, coord{y, x})
			}

		}
	}

	sum := 0
	for _, gear := range gearCoords {
		i := 0
		ratio := 1
		for _, num := range numbers {
			if hasAdj(num.coords, []coord{gear}) {
				i += 1
				ratio *= num.num
			}
		}
		if i == 2 {
			fmt.Println("gear", gear, i)
			sum += ratio
		}
	}

	return strconv.Itoa(sum)
}
