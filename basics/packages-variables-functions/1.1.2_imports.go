/*
	Notes 1.1.2

	- You can also import like this:
		import "fmt"
		import "math/rand"
*/

package main

import ( // this import is known as the 'factored import statement'
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
