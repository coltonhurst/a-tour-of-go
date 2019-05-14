/*
	Notes 1.2.2

	- The init and post statements are optional (condition is necessary)
*/

package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}