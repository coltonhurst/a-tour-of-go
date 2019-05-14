/*
	Notes 1.1.13

	- In Go, assignments of different types requires explicit
	  conversion, unlike C
	- To do a type conversion, use T(v) where it converts the value v
	  to type T
*/

package main

import (
	"fmt"
	"math"
)

// given sides of length 3 and 4 of a triangle, we will
// compute the hypotenuse sqrt(a^2 + b^2) because:
// c^2 = a^2 + b^2
func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
