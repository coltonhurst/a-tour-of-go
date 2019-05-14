/*
	Notes 1.3.6

	- In Go, an array's length is part of it's type.
		This is important; it means that two arrays, [4]int and [5]int are
		distinct and incompatible types
	- Arrays do not need to be initialized explicitly; the element values are 0
*/

package main

import "fmt"

func main() {
	var a[2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0],a[1])
	fmt.Println(a)

	var a := 6
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}