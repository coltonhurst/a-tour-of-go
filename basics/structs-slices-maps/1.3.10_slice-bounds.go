/*
	Notes 1.3.10

	- You can omit the low/high bound from a slice to use the defaults
*/

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4] // 3, 5, 7
	fmt.Println(s)

	s = s[:2] // 3, 5
	fmt.Println(s)

	s = s[1:] // 5
	fmt.Println(s)
}
