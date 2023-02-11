// A Go string is a read-only slice of bytes
// Concept of a character is called a `rune`
// it is an integer that represents a Unicode code point

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// `s` is a string assigned a literal value representing the 
	// world "hello" in Thai. Go string literals are UTF-0 encoded text
	const s = "สวัสดี"

	// since strings are equivalent to []byte, this will produce the
	// length of the raw bytes stored within
	fmt.Println("Len: ", len(s)) // 18
	
	// indexing into a string produces the raw byte values at each index
	// this loop generates the hex values of all the bytes that constitute the
	// code points in s
	for i:=0; i<len(s); i++ {
		fmt.Printf("%x ", s[i])
		// e0 b8 aa e0 b8 a7 e0 b8 b1 e0 b8 aa e0 b8 94 e0 b8 b5 %
	}

	fmt.Println()

	// to count how many `runes` are in a string
	// we can use the `utf8` package.
	
	fmt.Println("Rune Count: ", utf8.RuneCountInString(s))
	// Rune Count: 6
	
	// NOTE: run-time of `RuneCountInString` depends on the size of the string
	// because it has to decode each UTF-8 rune sequential

}