package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := []string{}
	findParenthesis(n, n, "", &res)
	return res
}

func findParenthesis(xindex, yindex int, str string, res *[]string) {
	if xindex == 0 && yindex == 0 {
		*res = append(*res, str)
		return
	}

	if xindex > 0 {
		findParenthesis(xindex-1, yindex, str+"(", res)
	}
	if yindex > 0 && xindex < yindex {
		findParenthesis(xindex, yindex-1, str+")", res)
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
