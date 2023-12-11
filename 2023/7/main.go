package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/sek-consulting/aoc/lib"
)

func main() {
	// lib.Expect(part1("example.txt"), "6440")
	// fmt.Println(part1("input.txt"))

	// lib.Expect(part2("example.txt"), "5905")
	fmt.Println(part2("input.txt"))
}

const (
	Strength1 string = "23456789TJQKA"
	Strength2 string = "J23456789TQKA"
)

type Type int

const (
	HighCard Type = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type hand struct {
	cards     string
	bet       int
	typ       Type
	strengths []int
}

func checkHand(hand string) Type {
	var counts []int
	hand2 := sortString(hand)
	for j := 0; j < len(hand2); j++ {
		if j == 0 || hand2[j-1] != hand2[j] {
			counts = append(counts, 1)
		} else {
			counts[len(counts)-1] += 1
		}
	}
	sort.SliceStable(counts, func(i, j int) bool { return counts[i] > counts[j] })

	// fmt.Println(hand, counts)

	switch counts[0] {
	case 5:
		return FiveOfAKind
	case 4:
		return FourOfAKind
	case 3:
		if counts[1] == 2 {
			return FullHouse
		}
		return ThreeOfAKind
	case 2:
		if counts[1] == 2 {
			return TwoPair
		}
		return OnePair
	default:
		return HighCard
	}
}

func compareHands(a, b Type) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func getStrengths(hand, strength string) []int {
	var strengths []int
	for _, c := range strings.Split(hand, "") {
		strengths = append(strengths, strings.Index(strength, c))
	}
	return strengths
}

func compareStrengths(a, b []int) int {
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			return 1
		} else if a[i] < b[i] {
			return -1
		}
	}
	return 0
}

func part1(filename string) string {
	var hands []hand
	for _, line := range lib.ReadLines(filename) {
		s := strings.Split(line, " ")

		hands = append(hands, hand{
			cards:     s[0],
			bet:       lib.Atoi(s[1]),
			typ:       checkHand(s[0]),
			strengths: getStrengths(s[0], Strength1),
		})

	}
	// fmt.Println("initial", hands)

	sort.SliceStable(hands, func(i int, j int) bool {
		typComp := compareHands(hands[i].typ, hands[j].typ)
		strengthComp := compareStrengths(hands[i].strengths, hands[j].strengths)
		return typComp < 0 || (typComp == 0 && strengthComp < 0)
	})
	// fmt.Println("sorted", hands)

	result := 0
	for i := 0; i < len(hands); i++ {
		// fmt.Println(i+1, hands[i].bet)
		result += hands[i].bet * (i + 1)
	}

	return strconv.Itoa(result)
}

func checkHand2(hand string) Type {
	var jokerCount int
	var counts []int
	hand2 := sortString(hand)
	for j := 0; j < len(hand2); j++ {
		if string(hand2[j]) == "J" {
			jokerCount += 1
		} else if j == 0 || hand2[j-1] != hand2[j] {
			counts = append(counts, 1)
		} else {
			counts[len(counts)-1] += 1
		}
	}
	sort.SliceStable(counts, func(i, j int) bool { return counts[i] > counts[j] })

	if len(counts) == 0 {
		counts = append(counts, jokerCount)
	} else {
		counts[0] += jokerCount
	}

	switch counts[0] {
	case 5:
		return FiveOfAKind
	case 4:
		return FourOfAKind
	case 3:
		if counts[1] == 2 {
			return FullHouse
		}
		return ThreeOfAKind
	case 2:
		if counts[1] == 2 {
			return TwoPair
		}
		return OnePair
	default:
		return HighCard
	}
}

func part2(filename string) string {
	var hands []hand
	for _, line := range lib.ReadLines(filename) {
		s := strings.Split(line, " ")

		hands = append(hands, hand{
			cards:     s[0],
			bet:       lib.Atoi(s[1]),
			typ:       checkHand2(s[0]),
			strengths: getStrengths(s[0], Strength2),
		})

	}
	fmt.Println("initial", hands)

	sort.SliceStable(hands, func(i int, j int) bool {
		typComp := compareHands(hands[i].typ, hands[j].typ)
		strengthComp := compareStrengths(hands[i].strengths, hands[j].strengths)
		return typComp < 0 || (typComp == 0 && strengthComp < 0)
	})
	fmt.Println("sorted", hands)

	result := 0
	for i := 0; i < len(hands); i++ {
		// fmt.Println(i+1, hands[i].bet)
		result += hands[i].bet * (i + 1)
	}

	return strconv.Itoa(result)
}
