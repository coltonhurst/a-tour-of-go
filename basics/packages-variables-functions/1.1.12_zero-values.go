/*
	Notes 1.1.12

	- Variables declared without an explicit initial value
	  are given their 'zero' value
	- Zero values:
	  - numeric types = 0
	  - boolean type = false
	  - string types = ""
	- Printf verbs:
	  - %v is the value in a default format
	  - %q a single quoted character literal that's safely escaped with Go syntax
	  - more: https://golang.org/pkg/fmt/
*/

package main

import (
	"fmt"
)

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
