/*
	Notes 1.1.9

	- a 'var' declaration can include initializers
	- If an initializer is present, you can leave out the type,
	  as the type will be inferred from the initializer
*/

package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
