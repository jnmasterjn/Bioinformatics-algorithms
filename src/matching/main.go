package main

import (
	"fmt"
)

func main() {
	fmt.Println("symbol matching")
}

// input 2 strings
// output the greatest number of matched symbols in any alignment of to strings
func matchingSymbol(str1, str2 string) int {
	n, m := len(str1), len(str2)

	lcs := make([][]int, n+1)

}
