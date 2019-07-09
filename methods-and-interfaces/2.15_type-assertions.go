/*
	Notes 2.15

	A type assertion provides access to an interface
	value's underlying concrete value.

	t := i.(T)
	
	The statement above will return the underlying T value to the variabel t,
	provided the interface value i holds the concrete type T. If it doesn't hold
	the concrete type T, a panic occurs.

	t, ok := i.(T)

	For this line, it's the same as above it; except 'ok' is a boolean.
	'ok' holds if the underlying concrete type is type T. If it doesn't hold
	the concrete type T, a panic will not occur, as 'ok' will just be false.
*/

package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}