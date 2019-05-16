/*
	Notes 2.3

	- You can declare a method on non-struct types also.
	- You can only declare a method with a receiver whose type is defined
	  in the same package as the method. So, because MyFloat is defined in
	  the 'main' package, we can make a method with a MyFloat as a receiver.
	  However, we can't make a method on a built in type like int, as it's
	  defined in a different package.

	Below, we create a numeric type MyFloat & give it an Abs() method.
*/

package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)	// math.Sqrt2 is a constant in the math package
	fmt.Println(f)				// before Abs()
	fmt.Println(f.Abs())		// after Abs()
}