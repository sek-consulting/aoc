package lib

import (
	"math"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func Atoi(s string) int {
	return Must(strconv.Atoi(s))
}

func IsNum(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func Abs[T constraints.Integer | constraints.Float](v T) T {
	if v < 0 {
		return -v
	}
	return v
}

func Pow(num, pow int) int {
	return int(math.Pow(float64(num), float64(pow)))
}

func Map[T, V any](items []T, mapper func(T) V) []V {
	result := make([]V, len(items))
	for i, item := range items {
		result[i] = mapper(item)
	}
	return result
}

func Filter[T any](items []T, filter func(T) bool) []T {
	var result []T
	for _, item := range items {
		if filter(item) {
			result = append(result, item)
		}
	}
	return result
}

func Every[T any](items []T, filter func(T) bool) bool {
	for _, item := range items {
		if !filter(item) {
			return false
		}
	}
	return true
}

func Reverse[T any](items []T) []T {
	result := make([]T, len(items))
	for i, item := range items {
		result[len(items)-i-1] = item
	}
	return result
}

func SumSlice[T constraints.Integer | constraints.Float](items []T) T {
	var sum T = 0
	for _, val := range items {
		sum += val
	}
	return sum
}

func MulSlice[T constraints.Integer | constraints.Float](items []T) T {
	var sum T = 1
	for _, val := range items {
		sum *= val
	}
	return sum
}

func MinSlice[T constraints.Ordered](items []T) T {
	var min T
	for i, val := range items {
		if i == 0 || val < min {
			min = val
		}
	}
	return min
}

func MaxSlice[T constraints.Ordered](items []T) T {
	var max T
	for i, val := range items {
		if i == 0 || val > max {
			max = val
		}
	}
	return max
}

func ReadLines(filename string) []string {
	data := Must(os.ReadFile(filename))
	return strings.Split(strings.TrimSpace(string(data)), "\r\n")
}
