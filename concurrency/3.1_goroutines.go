/*
	Notes 3.1

	A 'goroutine' is a "lightweight thread managed by the Go runtime".

	If you run go f(x, y, z), it evaluates f, x, y, and z on the current goroutine, and
	the execution of f happens in the new goroutine.
*/

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}