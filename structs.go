// structs are typed collections of fields
// useful for grouping data together to form records

package main

import "fmt"

// person struct type has name and age fields
type person struct {
	name string
	age  int
}

// newPerson constructs a new person struct with the
// given name.
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	// we can safely rturn a pointer to local variable
	// as local var will survive the scope of the function
	return &p
}

func main() {
	// syntax creates a new struct
	fmt.Println(person{"Bob", 20})

	// we can name the fields when initializing a struct
	fmt.Println(person{name: "Alice", age: 30})

	// omitted fields will be zero-valued
	fmt.Println(person{name: "John"})
	fmt.Println(person{age: 22})

	// An & prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 23})

	// it's idiomatic(natural) to encapsulate new struct creation in
	// constructor functions
	fmt.Println(newPerson("Jon"))

	// access struct fields with a dot
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age) // 50

	// structs are mutable
	sp.age = 51
	fmt.Println(sp.age) // 51
	fmt.Println(s.age)  // 51

}
