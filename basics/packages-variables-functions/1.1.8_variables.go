/*
	Notes 1.1.8

	- var can declare variables or a list of variables
	- like function parameters, the type goes on the end
*/

package main

import "fmt"

var c, python, java bool // create 3 booleans

func main() {
	var i int // create an int named i
	fmt.Println(i, c, python, java)
}
