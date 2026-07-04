package main

import "fmt"

// range is a Go keyword used to iterate over collections, returning the index/key and value of each element.
// iterating over data structures
func main() {
	nums := []int{6, 7, 8}

	for i, num := range nums {
		fmt.Println("first", num, i)
	}

	m := map[string]string{"fname": "john", "lname": "doe"}

	for k, v := range m {
		fmt.Println("second", k, v)
	}

	for k := range m {
		fmt.Println("third", k)
	}

	// unicode code point rune
	// starting byte of rune
	// 300 -> 1 byte , 2 byte
	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}

}
