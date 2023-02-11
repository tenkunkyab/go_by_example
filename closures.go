// go supports anon functions - which can form closures
// anon funcs are useful when we want to define a func inline (no name)

package main

import "fmt"

// `intSeq` returns another function
// which we define anon in the body of the `intSeq`
func intSeq() func() int {
	i := 0
	// returned func closes over `i` to form a closure
	return func() int {
		i++
		return i
	}
}

func main() {

	// we call `intSeq` assigning the result (func) to nextInt
	// this func value captures its own `i` value
	// which will be updated each time we call `nextInt`
	nextInt := intSeq()

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	newInts := intSeq()
	fmt.Println(newInts()) // 1
}
