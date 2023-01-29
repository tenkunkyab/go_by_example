package main

import "fmt"

// function that take two int and returns their
// sum as an int
func plus(a int, b int) int {
	// go requires explicit returns
	return a + b
}

// functions with multiple consequtive parameters of same
// type, we may omit the type name for the like-typed parameters
// up to the final parameter that declares the type
func multipleAdd(a, b, c int, d string) string {
	sum := a + b + c
	response := fmt.Sprintf("%s -> %d", d, sum)
	return response
}

func main() {
	res := plus(20, 22)
	fmt.Println("20 + 22: ", res)
	// 20 + 22:  42

	total := multipleAdd(21, 22, 23, "look at the total")
	fmt.Println(total)
	// look at the total -> 66
}
