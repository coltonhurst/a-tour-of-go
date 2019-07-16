/*
	Notes 3.2

	You can use channels to send and receive values with the channel operator, <-

	'chan'

	In the example below, we sum the numbers in the slice s. The work is
	distributed through two go routines (lines 29 and 30). When both goroutines
	complete their jobs, the final result is calculated.
*/

package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}