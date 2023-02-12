// Go supports `methods“ defined on struct types
package main

import "fmt"

type rect struct {
	width, height int
}

// area `method` has a receiver type of *rect
func (r *rect) area() int {
	return r.width * r.height
}

// `methods` can be defined for either pointer or
// value received types
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// here we call the two methods defined for our struct
	fmt.Println("area: ", r.area())   // 50
	fmt.Println("perim: ", r.perim()) // 30

	// go automatically handles conversions between values & pointers
	// for method calls. You may want to use a pointer receiver
	// type to avoid copying on method calls or to allow the method
	// to mutate the receiving struct
	rp := &r
	fmt.Println("area: ", rp.area())   // 50
	fmt.Println("perim: ", rp.perim()) // 30
}
