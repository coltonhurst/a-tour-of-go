/*
	Notes 2.14

	empty interface- the interface type that specifies zero methods

	Typed as: interface{}

	Empty interfaces can hold values of any type- each type implements at least zero methods.
	Empty interfaces are used by code that handles values of unkown type.
*/

package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}