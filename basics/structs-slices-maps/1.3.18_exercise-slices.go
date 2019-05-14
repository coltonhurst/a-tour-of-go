/*
	Notes 1.3.18

	For this one, copy and run the code on the 'A Tour of Go' website.
	This code will work and run, but you won't "see" the image.

	Remember, for a range, two values are returned:
	  - the index
	  - a copy of the element at that index

	You can omit something being returned with an underscore '_'. (second example)
	You can also omit the copy if you only want the index. (first example)
*/

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	// make the slice of uint8 slices
	// of length dy
	result := make([][]uint8, dy)

	// range through result, and set each 1st dimension
	// value to be a uint8 slice
	for i := range result {
		result[i] = make([]uint8, dx)
	}

	// range through each value of result, and set
	// the value based on where we are in the range
	for i := range result {
		for j := range result[i] {
			result[i][j] = uint8(0) + uint8(j)
		}
	}

	return result
}

func main() {
	pic.Show(Pic)
}
