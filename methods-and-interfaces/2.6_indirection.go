/*
	Notes 2.6

	Functions with a pointer argument take a pointer...
	Methods with pointer receivers take either a value or a pointer
	as the receiver when they are called.

	v.Scale(2), at the bottom, uses pointer indirection. Even though
	v is a value and not a pointer, the method with a pointer receiver is called automatically.

	So, v.Scale(2) is really (&v).Scale(2), since the Scale method has a pointer receiver.
*/

package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// a method with 
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}