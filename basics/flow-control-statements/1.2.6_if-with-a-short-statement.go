/*
	Notes 1.2.6

	- Just like 'for', 'if' can start with a short statement to execute before the condition
	- Any variables declared in this statement are out of scope at the end of the 'if'
*/

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
