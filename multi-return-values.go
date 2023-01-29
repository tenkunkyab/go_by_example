// go has built-in eupport for multiple return values
// this feature is used often in idiomatic Go
// for e.g. to return both result and error values from a func
package main

import "fmt"

// The (int, int) in this function signature shows
// that the function returns 2 ints
func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	// we use 2 different returnvalues from the call
	// with multiple assignment
	fmt.Printf("%d %d\n", a, b)
	// 3 7

	// if you only want a subset of the returned value
	// use the blank identifier `_`
	_, c := vals()
	fmt.Println(c)
	// 7
}
