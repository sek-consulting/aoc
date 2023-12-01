package lib

import (
	"os"
	"strconv"
	"strings"
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

func ReadLines(filename string) []string {
	data := Must(os.ReadFile(filename))
	return strings.Split(strings.TrimSpace(string(data)), "\r\n")
}
