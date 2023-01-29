// Variadic functions can be called with any number of
// trailing arguments. For e.g. `fmt.Println` is a common
// variadic function
package main

import "fmt"

// a function that will take an arbitrary number of
// ints as arguments
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Printf("Total: %d \n", total)
}

func main() {
	sum(1, 2)
	// [1 2]
	// Total: 3

	sum(1, 2, 3, 4)
	// [1 2 3 4]
	// Total: 10

	// If you already have multiple args in a slice, apply them to 
	// a variadic function using func(slice...) like this
	nums := []int{5, 6, 7, 8}
	sum(nums...)
	// 	[5 6 7 8]
	//  Total: 26

}
