package main

import "fmt"

// `zeroval` & `zeroptr.zeroval` has an int param
// so arguments will be passed to it by value

// `zeroval` will get a copy of `ival` distinct from the
// one in the calling function
func zeroval(ival int) {
	ival = 0
}

// `zeroptr` in contrast has an `*int` parameter
// it takes an `int` pointer.
// `*iptr` code in the func then dereferences the pointer
// from its memory address to the current value at that address
// ** Assigning a value to a dereferenced pointer changes the
// value at the referenced address**
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial: ", i) // 1
	zeroval(i)
	fmt.Println("zeroval: ", i) // 1
	// `&i` gives the memory address of `i` (pointer to `i`)
	zeroptr(&i)
	fmt.Println("zeroptr: ", i) // 0

	fmt.Println("pointer: ", &i) // 0xc000180b8
}
