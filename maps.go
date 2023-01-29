package main

import "fmt"

func main() {
	// Maps are Go's built-in associative data type (hashes, dicts)

	// make(map[key-type]val-type) creates an empty map
	m := make(map[string]int)

	// set key/value pairs using `name[key] = val` syntax
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m)
	// map:  map[k1:7 k2:13]

	// get a value for a key with name[key]
	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	// v1:  7

	fmt.Println("len: ", len(m))
	// len:  2

	// buildin `delete` removes key/value pairs from a map
	delete(m, "k2")
	fmt.Println("map: ", m)
	// map:  map[k1:7]

	// optional second return value when getting a value 
	// from a ap indicates if the `key` was present in the map
	// this can be used to disamiguate between missing keys and 
	// keys with zero values like `0` or `""`
	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	// prs: false


	// we can also declare and initialize a new map in the same
	// line with this syntax
	n:= map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)
	// map:  map[bar:2 foo:1]
	
}