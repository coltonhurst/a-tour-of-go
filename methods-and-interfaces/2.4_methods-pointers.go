/*
	Notes 2.4

	You can declare methods with pointer receivers.
	This allows you to change the value of the receiver passed by reference.

	Without the pointer, it's only passed by value, and any operations done
	are operating on a copy of the original, so after the method ends, the changes
	aren't saved.
*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}