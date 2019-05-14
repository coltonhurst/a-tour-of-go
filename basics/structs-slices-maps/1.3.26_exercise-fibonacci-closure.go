/*
	Notes 1.3.26

	In this exercise, we implement a fibonacci function that returns a closure that
	returns successive fibonacci numbres.

	fibonacciWithSlices & fibonacciSimple are functions that return
	a function that returns an int.

	0, 1, 1, 2, 3, 5 ...
*/

package main

import "fmt"

func fibonacciWithSlices() func() int {
	var nums []int = make([]int, 0, 0)

	return func() int {
		var toReturn int

		switch len(nums) {
		case 0:
			nums = append(nums, 0)
			toReturn =  nums[0]
		case 1:
			nums = append(nums, 1)
			toReturn =  nums[1]
		default:
			nums = append(nums, nums[len(nums) - 2] + nums[len(nums) - 1])
			toReturn = nums[len(nums) - 1]
		}

		return toReturn
	}
}

func fibonacciSimple() func() int {
	n := 0
	a := 0
	b := 1
	c := 0
	
	return func() int {
		if n == 0 {
			n++
			return 0
		} else if n == 1 {
			n++
			return 1
		} else {
			c = a + b
			a = b
			b = c
			return c
		}
	}
}

func main() {
	f := fibonacciSimple()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}