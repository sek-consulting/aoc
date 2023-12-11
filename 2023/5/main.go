package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sek-consulting/aoc/lib"
)

func main() {
	// lib.Expect(part1("example.txt"), "35")
	// fmt.Println(part1("input.txt"))

	// lib.Expect(part2("example.txt"), "46")
	fmt.Println(part2("input.txt"))
}

type mapping struct {
	start int
	end   int
	diff  int
}

func part1(filename string) string {
	var seeds []int
	var maps [][]mapping
	var curMaps []mapping

	isMapping := false
	for _, line := range lib.ReadLines(filename) {

		if strings.HasPrefix(line, "seeds: ") {
			s := strings.Split(line, "seeds: ")
			s = strings.Split(s[1], " ")
			for _, c := range s {
				seeds = append(seeds, lib.Atoi(c))
			}
		} else if strings.HasSuffix(line, " map:") {
			isMapping = true
			curMaps = make([]mapping, 0)
		} else if line == "" && isMapping {
			isMapping = false
			maps = append(maps, curMaps)
		} else if isMapping {

			values := lib.Map(strings.Split(line, " "), lib.Atoi)

			mapping := mapping{
				start: values[1],
				end:   values[1] + values[2],
				diff:  values[0] - values[1],
			}
			curMaps = append(curMaps, mapping)
		}
	}
	if isMapping {
		maps = append(maps, curMaps)
	}

	// fmt.Println(seeds)
	for i := 0; i < len(seeds); i++ {
		// fmt.Println("seed", seeds[i])
		for _, point := range maps {
			for _, mapping := range point {
				if lib.IsBetween(seeds[i], mapping.start, mapping.end) {
					seeds[i] += mapping.diff
					break
				}
			}
			// fmt.Println("->", seeds[i])
		}
	}
	// fmt.Println(seeds)

	min := lib.MinSlice(seeds)

	return strconv.Itoa(min)
}

type seed struct {
	start  int
	amount int
}

func part2(filename string) string {
	var seeds []seed
	var maps [][]mapping
	var curMaps []mapping

	isMapping := false
	for _, line := range lib.ReadLines(filename) {

		if strings.HasPrefix(line, "seeds: ") {
			s := strings.Split(line, "seeds: ")
			s = strings.Split(s[1], " ")
			for i := 0; i < len(s)-1; i += 2 {
				seeds = append(seeds, seed{lib.Atoi(s[i]), lib.Atoi(s[i+1])})
			}
		} else if strings.HasSuffix(line, " map:") {
			isMapping = true
			curMaps = make([]mapping, 0)
		} else if line == "" && isMapping {
			isMapping = false
			maps = append(maps, curMaps)
		} else if isMapping {

			values := lib.Map(strings.Split(line, " "), lib.Atoi)

			mapping := mapping{
				//mapping the other way around
				start: values[0],
				end:   values[0] + values[2],
				diff:  values[1] - values[0],
			}
			curMaps = append(curMaps, mapping)
		}
	}
	if isMapping {
		maps = append(maps, curMaps)
	}
	// fmt.Println(seeds)
	// fmt.Println(maps)

	solution := 0
	for i := 0; ; i++ {
		seed := i
		// fmt.Println(seed)
		for m := len(maps) - 1; m >= 0; m-- {
			for _, mapping := range maps[m] {
				if lib.IsBetween(seed, mapping.start, mapping.end) {
					seed += mapping.diff
					break
				}
			}
			// fmt.Println("->", seed)
		}
		if contains(seeds, seed) {
			solution = i
			break
		}
	}

	return strconv.Itoa(solution)
}

func contains(arr []seed, seed int) bool {
	for _, mapping := range arr {
		if lib.IsBetween(seed, mapping.start, mapping.start+mapping.amount) {
			return true
		}
	}
	return false
}
