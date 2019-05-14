/*
	Notes 1.2.13

	- Deferred function calls are pushed onto a stack
	- When the surrounding function returns, the deferred function calls
	  are executed in last-in-first-out order. (B/c it's a stack)
*/

package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
