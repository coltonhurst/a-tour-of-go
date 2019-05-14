/*
	Notes 1.1.11

	- Go's basic types:
	  - bool
	  - string
	  - int, int8, int16, int32, int64
	  - uint, uint8, uint16, uint32, uint64, uintptr
	  - byte (alias for uint8)
	  - rune (alias for int32, represents a unicode code point)
	  - float32, float64
	  - complex64, complex128
	- int, uint, and uintptr are 32bit / 64bit on a 32 bit / 64 bit system.
	  They recommend you stick with just int unless there are specific examples
*/

package main

import (
	"fmt"
	"math/cmplx"
)

// variable declarations can be "factored" into blocks like import statements
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type %T Value: %v\n", z, z)
}
