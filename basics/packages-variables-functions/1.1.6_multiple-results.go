/*
	Notes 1.1.6

	- Functions can return more than one result
	- ':=' is the short assignment statement. Used in place of a var declaration
	  with implicit type
*/

package main

import "fmt"

func swap(x, y string) (string, string) { // returning 2 strings
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
