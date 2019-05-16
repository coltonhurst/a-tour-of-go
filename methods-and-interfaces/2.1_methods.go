/*
	Notes 2.1

	- Go doesn't have classes
	- You define methods on types
	- A method is a function with a special receiver argument
	- The receiver appears in its own argument list between func and the method name
	- The Abs() method has a receiver of type Vertex named v
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

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}