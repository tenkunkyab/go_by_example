package main

import "fmt"

func main() {
	// range iterates over elements in a variety of data structures

	// we use `range` to sum the numbers in a slice
	// arrays work like this too
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	// sum: 9

	// ranges on arrays and slices provides both the index & value
	// for each entry. Since we don't need the index we used `_` to ignore
	// in case we need the index
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}
	// index: 1

	// range on map iterates over key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// a -> apple
	// b -> banana
	for k, v := range kvs {
		fmt.Println(k, "->", v)
	}
	// a -> apple
	// b -> banana

	// range can also iterate over just the keys of a map
	for k := range kvs {
		fmt.Println("key:", k)
	}
	// key: a
	// key: b

	// range on strings iterate over Unicode code points.
	// the first value is the starting byte index of the `rune`
	// and the second the `rune` itself.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
	// 0 103
	// 1 111
}
