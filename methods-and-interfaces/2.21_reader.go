/*
	Notes 2.21

	The io package provides the io.Reader interface.

	"Read populates the given byte slice with data and
	 returns the number of bytes populated and an error value. It returns an io.EOF error when the stream ends."
*/

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break;
		}
	}
}