/*
	Notes 1.1.7

	- You can name return values. If you do this, they are
	  treated as variables defined at the top of the function.
	- A return statement without arguments returns the named values.
	  This is known as a naked return.
	- Make sure you have good names, and only use naked returns in
	  small functions to ensure readability.
*/

package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // the naked return
}

func main() {
	fmt.Println(split(17))
}
