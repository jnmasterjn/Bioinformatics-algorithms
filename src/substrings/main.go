package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("sub string")

	s := "hi cmu"
	fmt.Println(s[0:5]) //start from 0 and end before 3 (only go to 2)

	pattern := "TA"
	text := "TATATATA"

	fmt.Println(PatternCount(pattern, text))
	//build in function in "string" package
	fmt.Println(strings.Count(text, pattern))

}

//takes input, two strings
//return numbers of time pattern occurs

func PatternCount(pattern, text string) int {
	count := 0
	k := len(pattern)
	n := len(text)

	//return over the text and find pattern
	for i := 0; i < n-k+1; i++ {
		if pattern == text[i:i+k] {
			count += 1
		}
	}
	return count
}

func StartIndices() {

}
