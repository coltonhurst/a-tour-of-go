/*
	Notes 1.1.5

	- When 2+ function params share the same type, you only need to put the type
	  after the last param
*/

package main

import "fmt"

func add(x, y int) int { // this, instead of: func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
