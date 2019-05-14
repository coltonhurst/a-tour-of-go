/*
	Notes 1.3.22

	When working with map 'm'...
	- Insert or update elements: m[key] = elem
	- Retrieving an element: elem = m[key]
	- Deleting an element: delete(m, key)
	- Testing that a key is present: elem, ok = m[key]
	  - If the key is in m, ok is true. If not, ok is false.
	  - If key is not in the map, then elem is the zero value for the map's element type.
	  - If they (elem or ok) haven't been declared yet... elem, ok := m[key]
*/

package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
