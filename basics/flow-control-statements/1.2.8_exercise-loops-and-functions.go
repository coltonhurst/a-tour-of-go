/*
	Notes 1.2.8

	- Implement a square root function with loops and functions
	- Using the variable name 'z', just following the AToG example
*/

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64 = 1
	for i := 0; i < 10; i++ {
		//fmt.Println(z)	// uncomment to see z become more precise
		z -= (z*z - x) / (2 * z)
	}
	return z
}

/*
	Here, we loop until the difference between the
	z and the previous z is less than 1E-10
*/
func SqrtAccurate(x float64) float64 {
	var z float64 = 1
	var z_previous float64 = 0
	var threshold float64 = 0.0000000001	// 1E-10

	for math.Abs(z - z_previous) >= threshold {
		//fmt.Println(z)	// uncomment to see z become more precise
		z_previous = z
		z -= (z*z - x) / (2 * z)
	}

	return z
}

func main() {
	fmt.Printf("Sqrt(%d) = %g\n", 2, Sqrt(2))
	fmt.Printf("SqrtAccurate(%d) = %g\n\n", 2, SqrtAccurate(2))

	fmt.Printf("Sqrt(%d) = %g\n", 2435, Sqrt(2435))
	fmt.Printf("SqrtAccurate(%d) = %g\n\n", 2435, SqrtAccurate(2435))

	fmt.Printf("Sqrt(%d) = %g\n", 25, Sqrt(25))
	fmt.Printf("SqrtAccurate(%d) = %g\n\n", 25, SqrtAccurate(25))
}