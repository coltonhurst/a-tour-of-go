/*
	Notes 1.3.17

	Remember, for a range, two values are returned:
	  - the index
	  - a copy of the element at that index

	You can omit something being returned with an underscore '_'. (second example)
	You can also omit the copy if you only want the index. (first example)
*/

package main

import "fmt"

func main() {
	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
