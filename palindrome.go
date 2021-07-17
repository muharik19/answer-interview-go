package main

import (
	"fmt"
	"strings"
)

func pd(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var input string
	fmt.Print("Type some palindrome: ")
	fmt.Scanln(&input)
	noSpaceString := strings.ReplaceAll(input, input, "")
	fmt.Println(pd(noSpaceString))
}
