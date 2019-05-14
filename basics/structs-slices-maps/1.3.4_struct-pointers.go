/*
	Notes 1.3.4

	- You can make struct pointers
	- We could write (*p).X to access the X field in the 'v' Vertex struct.
		Shorthand just uses the dot operator, removing the explicit dereference
	- You can obtain these from slice s with len(s) and cap(s)
*/

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}