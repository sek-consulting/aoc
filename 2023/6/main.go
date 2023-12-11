package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sek-consulting/aoc/lib"
)

func main() {
	// lib.Expect(part1("example.txt"), "288")
	// fmt.Println(part1("input.txt"))

	// lib.Expect(part2("example.txt"), "71503")
	fmt.Println(part2("input.txt"))
}

func part1(filename string) string {
	var times []int
	var distances []int

	for _, line := range lib.ReadLines(filename) {
		if strings.HasPrefix(line, "Time:") {
			for _, i := range strings.Split(line, " ") {
				s := strings.TrimSpace(i)
				if lib.IsNum(s) {
					times = append(times, lib.Atoi(s))
				}
			}
		} else if strings.HasPrefix(line, "Distance:") {
			for _, i := range strings.Split(line, " ") {
				s := strings.TrimSpace(i)
				if lib.IsNum(s) {
					distances = append(distances, lib.Atoi(s))
				}
			}
		}
	}
	fmt.Println(times, distances)

	solution := 1
	for i := 0; i < len(times); i++ {
		fmt.Println(times[i], "->", distances[i])
		min := 0
		for speed := 1; speed < times[i]; speed++ {
			if speed*(times[i]-speed) > distances[i] {
				fmt.Println("min", speed)
				min = speed
				break
			}
		}
		for speed := times[i]; speed > 0; speed-- {
			if speed*(times[i]-speed) > distances[i] {
				fmt.Println("min", speed)
				solution *= speed - min + 1
				break
			}
		}
	}

	return strconv.Itoa(solution)
}

func part2(filename string) string {
	var time int
	var distance int

	for _, line := range lib.ReadLines(filename) {
		if strings.HasPrefix(line, "Time:") {
			s := strings.Split(line, "Time:")[1]
			s = strings.ReplaceAll(s, " ", "")
			time = lib.Atoi(s)
		} else if strings.HasPrefix(line, "Distance:") {
			s := strings.Split(line, "Distance:")[1]
			s = strings.ReplaceAll(s, " ", "")
			distance = lib.Atoi(s)
		}
	}
	fmt.Println(time, distance)

	solution := 0
	for speed := 1; speed < time; speed++ {
		if speed*(time-speed) > distance {
			fmt.Println("min", speed)
			solution = speed
			break
		}
	}
	for speed := time; speed > 0; speed-- {
		if speed*(time-speed) > distance {
			fmt.Println("min", speed)
			solution = speed - solution + 1
			break
		}
	}

	return strconv.Itoa(solution)
}
