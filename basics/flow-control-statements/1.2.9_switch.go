/*
	Notes 1.2.9

	- Go switch statements 'automatically have break statements'
    (when a case matches, it's the only one that will be run)
*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Println("%s.", os)
	}
}
