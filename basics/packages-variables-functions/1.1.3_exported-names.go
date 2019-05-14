/*
	Notes 1.1.3

	- A name is exported if it begins with a capital letter
	- "unexported" names are not available outside the package
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi) // notice the capital 'P'
}
