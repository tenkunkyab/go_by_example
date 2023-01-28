package main

import "fmt"

func main() {
	// unlike arrays, slices are typed only by the elements
	// they contain (not the number of elements)

	// to create an empty slice with non-zero length, we use
	// ** make **
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s, s[2], len(s))
	// set:  [a b c] c 3

	// slices support several more operations compared to array
	// 1. append
	s = append(s, "d")
	fmt.Println("1st append: ", s)
	// 1st append:  [a b c d]
	x := append(s, "e", "f") // reference
	fmt.Println("2nd append: ", s, x)
	// 2nd append:  [a b c d] [a b c d e f]

	x[1] = "op"
	fmt.Println(s, x)
	// [a op c d] [a op c d e f]

	/**

	https://go.dev/blog/slices-intro

	slice can either be created
	1. literal
		letters := []string{"a", "b", "c"}
	2. built-in function `make`
		```
		func make([]T, len, cap) [] T
		````
		T = element type of slice to be created
		len = length
		cap = optional capacity, if not provided defaults to `len`

		when called `make` allocates an array and returns
		a slice that refers to that array

		```
		var s []byte
		s = make([]byte, 5, 5)
		s == []byte{0, 0, 0, 0, 0}
		```

		the `length` and `capacity` of a slice can be inspected using
		built-in `len` and `cap` function
		len(s)
		cap(s)
	**/

	// 2. copy'd -- we create an empty slice `c` of the same
	// length as `s` and copy into `c` from `s`
	c := make([]string, len(s))
	copy(c, s) // copy by value
	fmt.Println("cpy: ", c)
	// cpy:  [a op c d]

	s[0] = "ab"
	fmt.Println(s, x, c)
	// [ab op c d] [ab op c d e f] [a op c d]

	// 3. `slice` operation with `slice[low:high]`
	l := s[2:5] // copied by reference
	fmt.Println(`sl1: `, l)
	// sl1:  [c d e]
	fmt.Println(s, x, c)
	// [a op c d] [a op c d e f] [ab op c d]
	fmt.Println(s, len(s), s[0], s[1], s[2], s[3], s[2:6])
	// [ab op c d] 4 ab op c d [c d e f]
	// wait what? how can s[2:6] which slices
	// s from 2nd index till 5th index i.e. s[2], s[3], s[5]
	// when len(s) is 4! where is it picking up `e`, and `f` from?

	// answer:
	// in the memory
	// a, op, c, d, e, f
	// we have three slices s, x, c
	// s points to the location of `a`
	// len(s) = 4 hence when we print s => a, op, c, d
	// but cap(s) = 6 hence when we do s[2:6] => c, d, e, f
	// x points to the location of `a`
	// len(x) = 6 hence when we print x => a, op, c, d, e, f

	l[1] = "poop"
	fmt.Println(l)
	// [c poop e]
	fmt.Println(s, x, c)
	// [ab op c poop] [ab op c poop e f] [a op c d]

	t := []string{"g", "h", "i"}
	fmt.Println("dcl: ", t)

	// slices can be composed into multi-dimensional data structures
	// the length of inner slices can vary, unlike with
	// multi-dimensional arrays
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
