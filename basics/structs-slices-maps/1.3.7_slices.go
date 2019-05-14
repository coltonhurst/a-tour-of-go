/*
	Notes 1.3.6

	- A slice is a dynamically-sized, flexible view into the elements of an array
	- Slices are more common than arrays in practice
	- Specify a slice like this: a[low : high]
		It excludes the end (high)
	- The difference between arrays and slices when declared are subtle;
	  if there is a size specified, it's an array.
*/

package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13} // This is an array

	var s []int = primes[1:4] // This is a slice. Selects [3, 5, 7] from primes.
	fmt.Println(s)
}
