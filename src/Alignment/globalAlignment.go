package main

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
		if str1[roww-1] == str2[coll-1] {

		}

		//case 2, mismatch, have two different letters
		else if [roww-1] != str2[coll-1] {

		}




	}

}
