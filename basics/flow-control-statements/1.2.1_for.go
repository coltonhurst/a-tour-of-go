/*
	Notes 1.2.1

	- Go only has one looping construct- the for loop!
	- There are three components separated by semicolons:
	  - the init statement
	  - condition expression
	  - post statement
	- There are no parenthesis around the 3 components in the "loop header" (what I'm calling it)
	- Braces on a loop { } are always required
*/

package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}