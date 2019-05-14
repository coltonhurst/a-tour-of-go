/*
	Notes 1.1.14

	- In Go, if you don't specify a type explicitly when delcaring a variable,
	  the type will be inferred.
	- If the variable's initialization value is an untyped numeric constant,
	  the new variable may be an int, float64, or complex128 depending on the precision.
*/

package main

import "fmt"

func main() {
	v := 42 // change this (example values: 42, 3.142, 0.867 + 0.5i)
	fmt.Printf("v is of type %T\n", v)
}
