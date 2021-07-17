package main

import (
	"fmt"
)

func fb(n int) int {
	a := 0
	b := 1
	for i := 0; i < n; i++ {
		temp := a
		a = b
		b = temp + a
	}
	return a
}

func main() {
	var input int
	fmt.Print("Type some palindrome: ")
	fmt.Scanln(&input)

	for n := 0; n < input; n++ {
		result := fb(n)
		fmt.Printf("Fibonacci %v = %v\n", n, result)
	}
}
