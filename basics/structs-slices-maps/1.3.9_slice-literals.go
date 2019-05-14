/*
	Notes 1.3.9

	- A slice literal is like an array literal... without the length
	- Array literal: [3]bool{true, true, false}
	- Slice literal: []bool{true, true, false}
	  This will build an array, and then a slice that references it
*/

package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}