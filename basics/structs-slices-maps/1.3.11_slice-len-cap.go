/*
	Notes 1.3.11

	- A slice has a length and a capacity
		- length: number of elements in the slice
		- capacity: number of elements in the underlying array, counting from the
			first element of the slice
	- %v is the value in a default format (used in the printf())
*/

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlices(s)

	// Slice the slice to give it zero length
	s = s[:0]
	printSlices(s)

	// Extend its length.
	s = s[:4]
	printSlices(s)

	// Drop its first two values.
	s = s[2:]
	printSlices(s)
}

func printSlices(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
