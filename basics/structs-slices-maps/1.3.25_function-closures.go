/*
	Notes 1.3.25

	Go functions can be closures!

	Reminder- a closure is a function value that references variables from outside it's own function body.
	It may access and assign these referenced variables- so, the function is "bound" to the variables.

	Below, the adder() function returns a closure. Each closure is bound to its own sum variable.
*/

package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
