/*
	Notes 2.8

	Two reasons to use a pointer receiver
	1) so the method can modify the value the receiver points to
	2) to avoid copying the value on each method call

	Notice, in the Abs() method, it doesn't modify the receiver.
	However, b/c it's a pointer receiver method, we don't have to copy
	the value on each method call. (Could acheive the same with just passing a pointer.)

	"In general, all methods on a given type should have either value or pointer receivers,
	but not a mixture of both."
*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before caling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}