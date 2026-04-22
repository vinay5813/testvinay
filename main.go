package main

import "fmt"

func isEven(n int) bool {
	return n%2 == 0
}

func add(a, b int) int {
	return a + b
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
}
