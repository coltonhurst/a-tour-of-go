/*
	Notes 2.22

	Exercise: https://tour.golang.org/methods/22

	"Implement a Reader type that emits an infinite stream of the ASCII character 'A'."
*/

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 65
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}