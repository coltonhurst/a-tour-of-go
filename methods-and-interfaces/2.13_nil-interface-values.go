/*
	Notes 2.13

	A nil interface value holds neither a value nor a concrete type.
	Calling a method on a nil interface causes a runtime error, b/c
	there's no type in the interface tuple to indicate which concrete
	method to call.

	A comment also was written about this in Notes 2.12.
*/

package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}