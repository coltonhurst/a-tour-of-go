/*
	Notes 1.3.19

	A map "maps" keys to values. In Golang, maps are implemented as hash tables.

	The zero value of a map is nil.
*/

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
