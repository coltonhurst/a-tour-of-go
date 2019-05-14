/*
	Notes 1.1.4

	- Functions can take 0+ arguments
	- notice the type comes after the param name
	- return type is at the end of the function header
	- more info: https://blog.golang.org/gos-declaration-syntax
*/

package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
