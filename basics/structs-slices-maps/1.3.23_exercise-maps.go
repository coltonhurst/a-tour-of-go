/*
	Notes 1.3.23

	For this one, copy and run the code on the 'A Tour of Go' website.
	This code is made for the test suite built into this exercise:
	https://tour.golang.org/moretypes/23

	To be more efficient, the for loop / range could just contain this in the body:
	'wordCount[word]++'
*/

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)        // splits the string s into a string array, splitting by spaces, testing with 'unicode.IsSpace'
	wordCount := make(map[string]int) // expanded: var wordCount map[string]int = make(map[string]int)

	for _, word := range words { // range through each word in 'words'
		_, exists := wordCount[word] // wordCount is the map (the keys are strings, values are ints), get if the word exists in the map

		if exists {
			wordCount[word] = wordCount[word] + 1
		} else {
			wordCount[word] = 1
		}
	}

	return wordCount
}

func main() {
	wc.Test(WordCount)
}
