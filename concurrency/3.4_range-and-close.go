/*
	Notes 3.4

	A sender can close channels to indicate no more values will be sent.
	Receivers can test if channels are closed by testing the ok value:

	v, ok := <-ch

	ok is false if the channel is closed and there are no more values to receive.

	i := range c

	This loop receives values from the channel until it's closed.

	PS:
	- Only the sender should close a channel, never the receiver (sending on a closed channel)
	will cause a panic
	- Channels aren't like files; you don't *usually* need to close them. You only need to close them
	when the receiver must be told there are no more values coming (like when terminating a range loop).
*/

package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)	// buffered channel with a size of 10
	go fibonacci(cap(c), c)
	for v := range c {
		fmt.Println(v)
	}
}