/*
	Notes 1.1.10

	- The := short assignment statement can be used in a function
	  in place of a var declaration
	- := is not available outside a function (outside a func every
	  statement begins with a keyword)
*/

package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
