/*
	Notes 1.2.12

	- A defer statement 'defers' (delays) the execution of a function
	  until the surrounding function returns.
*/

package main

import "fmt"

func main() {
	defer fmt.Println("world!")

	fmt.Println("Hello ")
}