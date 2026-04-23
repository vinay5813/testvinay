package main

import (
	"fmt"
	"unicode"
)

func isEven(n int) bool {
	return n%2 == 0
}

func add(a, b int) int {
	return a + b
}

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

	numbers := []int{1, 2, 3, 4, 5}
	for _, n := range numbers {
		if isEven(n) {
			fmt.Printf("%d is even\n", n)
		} else {
			fmt.Printf("%d is odd\n", n)
		}
	}

	sum := add(10, 20)
	fmt.Printf("Sum of 10 and 20 = %d\n", sum)

	words := []string{"12345", "abc123", "hello", "!@#$"}
	for _, w := range words {
		fmt.Printf("%q is %s\n", w, checkString(w))
	}
}
