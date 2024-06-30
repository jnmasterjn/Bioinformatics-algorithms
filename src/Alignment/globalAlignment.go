package main

import (
	"fmt"
)

func main() {
	fmt.Println(GlobalAlignment("GATATAA", "GTATAGATA", 1, 2, 3))
}

// GlobalAlignment takes two strings, along with match, mismatch, and gap scores.
// It returns a maximum score global alignment of the strings corresponding to these penalties.

// Sample Input 1:

// 1 1 2
// GAGA
// GAT
// Sample Output 1:

// A sample correct alignment for this problem is:
// GAGA
// GAT-
// Your alignment should have the same score as the sample alignment: -1
// Sample Input 2:

// 1 3 1
// ACG
// ACT
// Sample Output 2:

// A sample correct alignment for this problem is:
// AC-G
// ACT-
// Your alignment should have the same score as the sample alignment: 0

type Alignment [2]string

func GlobalAlignment(str1, str2 string, match, mismatch, gap float64) Alignment {
	var a Alignment //arrary containing two strings

	table := GlobalScoreTable(str1, str2, match, mismatch, gap) //make the score table (matrix)

	//row and column
	roww := len(str1)
	coll := len(str2)

	//backtracking
	//loop when haven't reach top left
	for roww != 0 && coll != 0 {
		//case 1, both letter from the 2 strings match (diag)
		if str1[roww-1] == str2[coll-1] && table[roww-1][coll-1]+match == table[roww][coll] {
			//str add in the front, so don't need to reverse later
			a[0] = string(str1[roww-1]) + a[0] //first string in alignment, store matched value
			a[1] = string(str2[coll-1]) + a[1] //second string in alignment store matched value
			roww--
			coll-- //backtrack, row and col will decrease as we go on
		} else if table[roww-1][coll-1]-mismatch == table[roww][coll] { //case 2, mismatch, have two different letters
			a[0] = string(str1[roww-1]) + a[0]
			a[1] = string(str2[coll-1]) + a[1]
			roww--
			coll--
		} else if table[roww][coll-1]-gap == table[roww][coll] { //case 3, mismatch,left, missing letter for string 1, need to add "-"
			a[0] = "-" + a[0]
			a[1] = string(str2[coll-1]) + a[1]
			coll--
		} else if table[roww-1][coll]-gap == table[roww][coll] { //case 4, up,mismatch, missing letter for string 2
			a[0] = string(str1[roww-1]) + a[0]
			a[1] = "-" + a[1]
			roww--
		}

		for roww != 0 && coll == 0 {
			a[0] = string(str1[roww]) + a[0]
			a[1] = "-" + a[1]
			roww--
		}

		for coll != 0 && roww == 0 {
			a[0] = "-" + a[0]
			a[1] = string(str2[coll]) + a[1]
		}
	}
	return a
}
