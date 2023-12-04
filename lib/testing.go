package lib

import "fmt"

func Expect(result, expected string) {
	if result != expected {
		fmt.Printf("Result was incorrect, got: %s, want: %s.\n", result, expected)
	} else {
		fmt.Println("Success!")
	}
}
