/*
	Notes 1.1.1

	- programs start running in the package 'main'
	- the package name is the same as the last element of the import path
	- see 'rand.Seed' so the 'rand.Intn()' function doesn't always return the same value
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is ", rand.Intn(10))
}
