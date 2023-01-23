package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	// := syntax is shorthand for declaring and initializing a variable
	// same as var f string = "apple"
	f := "apple"
	fmt.Println(f)

	var g float32 = 123456789
	fmt.Println(g)

	var h float64 = 123456789123456789
	fmt.Println(h)
}
