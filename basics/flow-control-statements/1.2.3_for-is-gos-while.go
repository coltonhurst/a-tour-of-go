/*
	Notes 1.2.3

	- Well, you don't even need the semicolons!
	  If you just have a condition, it's just like a C while loop
*/

package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}