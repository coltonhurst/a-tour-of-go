/*
	Notes 1.3.16

	'range' is a form of a for loop that iterates over a slice or a map.

	Two values are always returned each iteration: the index, and the copy of the
	element at that index.
*/

package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}