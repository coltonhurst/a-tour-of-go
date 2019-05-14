/*
	Notes 1.3.1

	- Pointers hold the memory address of a value
	- Zero value is nil
	- The & operator gives the address of the variable
	- * dereferences the pointer
	- Go has no pointer arithmetic
*/

package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer p
	*p = 21         // set i through pointer p
	fmt.Println(i)  // see the new value of i

	p = &j         // point p to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
