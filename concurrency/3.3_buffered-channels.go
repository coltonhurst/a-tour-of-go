/*
	Notes 3.3

	Channels can be buffered- provide the length during channel creation with make().

	"Sends to a buffered channel block only when the buffer is full. Receives block
	when the buffer is empty."

	If you overfill the buffer, you'll get an error.
*/

package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}