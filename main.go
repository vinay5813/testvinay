package main

import (
	"fmt"
	"unicode"
)

// isEven returns true if n is an even number, false if odd.
func isEven(n int) bool {
	return n%2 == 0
}

// add returns the sum of two integers a and b.
func add(a, b int) int {
	return a + b
}

// checkString checks whether the given string is numeric, alphanumeric, or neither.
// Returns "numeric" if all characters are digits.
// Returns "alphanumeric" if all characters are letters or digits.
// Returns "neither" if the string contains special characters.
// Returns "empty" if the string has no characters.
func checkString(s string) string {
	if len(s) == 0 {
		return "empty"
	}
	isNum := true
	isAlnum := true
	for _, ch := range s {
		if !unicode.IsDigit(ch) {
			isNum = false
		}
		if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
			isAlnum = false
		}
	}
	if isNum {
		return "numeric"
	}
	if isAlnum {
		return "alphanumeric"
	}
	return "neither"
}

func main() {
	fmt.Println("hello vinay how are you")

	// Check even or odd for a list of numbers
	numbers := []int{1, 2, 3, 4, 5}
	for _, n := range numbers {
		if isEven(n) {
			fmt.Printf("%d is even\n", n)
		} else {
			fmt.Printf("%d is odd\n", n)
		}
	}

	// Add two numbers and print the result
	sum := add(10, 20)
	fmt.Printf("Sum of 10 and 20 = %d\n", sum)

	// Check each string and print whether it is numeric, alphanumeric, or neither
	words := []string{"12345", "abc123", "hello", "!@#$"}
	for _, w := range words {
		fmt.Printf("%q is %s\n", w, checkString(w))
	}
}
